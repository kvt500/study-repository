package main

import (
	"context"
	"study/feature_postgres/simple_connection"
	"study/simple_sql"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	// if err := simple_sql.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"Обед",
	// 	"Покущать надо",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.UpdateRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	pp.Println(tasks)
}
