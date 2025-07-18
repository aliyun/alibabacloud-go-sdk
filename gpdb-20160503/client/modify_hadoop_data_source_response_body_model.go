// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHadoopDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHadoopDataSourceResponseBody
	GetRequestId() *string
}

type ModifyHadoopDataSourceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 2C125605-266F-41CA-8AC5-3A643D4F42C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHadoopDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHadoopDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHadoopDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHadoopDataSourceResponseBody) SetRequestId(v string) *ModifyHadoopDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHadoopDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
