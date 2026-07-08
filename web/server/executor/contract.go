package executor

import "github.com/mailcache/goconvey/web/server/contract"

type Parser interface {
	Parse([]*contract.Package)
}

type Tester interface {
	SetBatchSize(batchSize int)
	TestAll(folders []*contract.Package)
}
