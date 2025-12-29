// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateScreenSaverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddOrUpdateScreenSaverResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddOrUpdateScreenSaverResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddOrUpdateScreenSaverResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *AddOrUpdateScreenSaverResponseBody
	GetStatusCode() *int32
}

type AddOrUpdateScreenSaverResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EED***9661
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateScreenSaverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateScreenSaverResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddOrUpdateScreenSaverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddOrUpdateScreenSaverResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddOrUpdateScreenSaverResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateScreenSaverResponseBody) SetMessage(v string) *AddOrUpdateScreenSaverResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponseBody) SetRequestId(v string) *AddOrUpdateScreenSaverResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponseBody) SetResult(v bool) *AddOrUpdateScreenSaverResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponseBody) SetStatusCode(v int32) *AddOrUpdateScreenSaverResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponseBody) Validate() error {
	return dara.Validate(s)
}
