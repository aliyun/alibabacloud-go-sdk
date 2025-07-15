// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int32) *CreateAppResponseBody
	GetAppId() *int32
	SetRequestId(v string) *CreateAppResponseBody
	GetRequestId() *string
}

type CreateAppResponseBody struct {
	// The ID of the application.
	//
	// example:
	//
	// 1234
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) GetAppId() *int32 {
	return s.AppId
}

func (s *CreateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppResponseBody) SetAppId(v int32) *CreateAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) Validate() error {
	return dara.Validate(s)
}
