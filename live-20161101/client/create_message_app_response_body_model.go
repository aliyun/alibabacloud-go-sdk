// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMessageAppResponseBody
	GetRequestId() *string
	SetResult(v *CreateMessageAppResponseBodyResult) *CreateMessageAppResponseBody
	GetResult() *CreateMessageAppResponseBodyResult
}

type CreateMessageAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *CreateMessageAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMessageAppResponseBody) GetResult() *CreateMessageAppResponseBodyResult {
	return s.Result
}

func (s *CreateMessageAppResponseBody) SetRequestId(v string) *CreateMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMessageAppResponseBody) SetResult(v *CreateMessageAppResponseBodyResult) *CreateMessageAppResponseBody {
	s.Result = v
	return s
}

func (s *CreateMessageAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMessageAppResponseBodyResult struct {
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateMessageAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMessageAppResponseBodyResult) GetAppId() *string {
	return s.AppId
}

func (s *CreateMessageAppResponseBodyResult) SetAppId(v string) *CreateMessageAppResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *CreateMessageAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
