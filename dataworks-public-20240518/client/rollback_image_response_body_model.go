// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RollbackImageResponseBody
	GetData() *bool
	SetRequestId(v string) *RollbackImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RollbackImageResponseBody
	GetSuccess() *bool
}

type RollbackImageResponseBody struct {
	// The result of the API request.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackImageResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *RollbackImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackImageResponseBody) SetData(v bool) *RollbackImageResponseBody {
	s.Data = &v
	return s
}

func (s *RollbackImageResponseBody) SetRequestId(v string) *RollbackImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackImageResponseBody) SetSuccess(v bool) *RollbackImageResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackImageResponseBody) Validate() error {
	return dara.Validate(s)
}
