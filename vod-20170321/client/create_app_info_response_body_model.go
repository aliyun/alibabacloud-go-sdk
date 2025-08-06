// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppInfoResponseBody
	GetAppId() *string
	SetRequestId(v string) *CreateAppInfoResponseBody
	GetRequestId() *string
}

type CreateAppInfoResponseBody struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-34D5-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInfoResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppInfoResponseBody) SetAppId(v string) *CreateAppInfoResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppInfoResponseBody) SetRequestId(v string) *CreateAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
