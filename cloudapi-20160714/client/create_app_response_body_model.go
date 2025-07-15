// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *CreateAppResponseBody
	GetAppId() *int64
	SetRequestId(v string) *CreateAppResponseBody
	GetRequestId() *string
	SetTagStatus(v bool) *CreateAppResponseBody
	GetTagStatus() *bool
}

type CreateAppResponseBody struct {
	// The unique ID of the application.
	//
	// example:
	//
	// 20112314518278
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BA20890E-75C7-41BC-9C8B-73276B58F550
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the tag exists. If the value is **true**, the tag exists. If the value is **false**, the tag does not exist.
	//
	// example:
	//
	// false
	TagStatus *bool `json:"TagStatus,omitempty" xml:"TagStatus,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) GetAppId() *int64 {
	return s.AppId
}

func (s *CreateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppResponseBody) GetTagStatus() *bool {
	return s.TagStatus
}

func (s *CreateAppResponseBody) SetAppId(v int64) *CreateAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetTagStatus(v bool) *CreateAppResponseBody {
	s.TagStatus = &v
	return s
}

func (s *CreateAppResponseBody) Validate() error {
	return dara.Validate(s)
}
