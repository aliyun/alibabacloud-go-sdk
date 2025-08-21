// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateApplicationResponseBody
	GetAppId() *string
	SetRequestId(v string) *CreateApplicationResponseBody
	GetRequestId() *string
}

type CreateApplicationResponseBody struct {
	// The ID of the application. This is the unique identifier of the edge application.
	//
	// example:
	//
	// d0639abf-789a-4527-b420-031d2cd9ad9b
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationResponseBody) SetAppId(v string) *CreateApplicationResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
