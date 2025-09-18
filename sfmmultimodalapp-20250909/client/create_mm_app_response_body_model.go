// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMmAppResponseBody
	GetAppId() *string
	SetRequestId(v string) *CreateMmAppResponseBody
	GetRequestId() *string
}

type CreateMmAppResponseBody struct {
	// example:
	//
	// mm-xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *CreateMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMmAppResponseBody) SetAppId(v string) *CreateMmAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateMmAppResponseBody) SetRequestId(v string) *CreateMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMmAppResponseBody) Validate() error {
	return dara.Validate(s)
}
