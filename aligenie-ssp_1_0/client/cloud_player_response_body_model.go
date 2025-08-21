// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPlayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CloudPlayerResponseBody
	GetCode() *int32
	SetMessage(v string) *CloudPlayerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudPlayerResponseBody
	GetRequestId() *string
	SetResult(v bool) *CloudPlayerResponseBody
	GetResult() *bool
}

type CloudPlayerResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B7D82DB0-DD59-5756-AC62-871C9D7DBC36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudPlayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudPlayerResponseBody) GoString() string {
	return s.String()
}

func (s *CloudPlayerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CloudPlayerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudPlayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudPlayerResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CloudPlayerResponseBody) SetCode(v int32) *CloudPlayerResponseBody {
	s.Code = &v
	return s
}

func (s *CloudPlayerResponseBody) SetMessage(v string) *CloudPlayerResponseBody {
	s.Message = &v
	return s
}

func (s *CloudPlayerResponseBody) SetRequestId(v string) *CloudPlayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudPlayerResponseBody) SetResult(v bool) *CloudPlayerResponseBody {
	s.Result = &v
	return s
}

func (s *CloudPlayerResponseBody) Validate() error {
	return dara.Validate(s)
}
