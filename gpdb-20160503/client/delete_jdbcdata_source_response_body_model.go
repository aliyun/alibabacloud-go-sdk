// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJDBCDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteJDBCDataSourceResponseBody
	GetRequestId() *string
}

type DeleteJDBCDataSourceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJDBCDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJDBCDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJDBCDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJDBCDataSourceResponseBody) SetRequestId(v string) *DeleteJDBCDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJDBCDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
