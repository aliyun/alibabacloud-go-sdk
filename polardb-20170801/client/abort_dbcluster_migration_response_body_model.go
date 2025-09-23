// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortDBClusterMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AbortDBClusterMigrationResponseBody
	GetRequestId() *string
}

type AbortDBClusterMigrationResponseBody struct {
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortDBClusterMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *AbortDBClusterMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortDBClusterMigrationResponseBody) SetRequestId(v string) *AbortDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortDBClusterMigrationResponseBody) Validate() error {
	return dara.Validate(s)
}
