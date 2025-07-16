// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveContentResponseBody
	GetRequestId() *string
	SetResult(v string) *SaveContentResponseBody
	GetResult() *string
	SetSuccess(v bool) *SaveContentResponseBody
	GetSuccess() *bool
}

type SaveContentResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveContentResponseBody) GoString() string {
	return s.String()
}

func (s *SaveContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveContentResponseBody) GetResult() *string {
	return s.Result
}

func (s *SaveContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveContentResponseBody) SetRequestId(v string) *SaveContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveContentResponseBody) SetResult(v string) *SaveContentResponseBody {
	s.Result = &v
	return s
}

func (s *SaveContentResponseBody) SetSuccess(v bool) *SaveContentResponseBody {
	s.Success = &v
	return s
}

func (s *SaveContentResponseBody) Validate() error {
	return dara.Validate(s)
}
