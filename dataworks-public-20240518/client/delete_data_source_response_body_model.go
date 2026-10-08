// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataSourceResponseBody
	GetSuccess() *bool
}

type DeleteDataSourceResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// B56432E0-2112-5C97-88D0-AA0AE5C75C74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call is successful.
	//
	// - true: Successful
	//
	// - false: Failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetSuccess(v bool) *DeleteDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
