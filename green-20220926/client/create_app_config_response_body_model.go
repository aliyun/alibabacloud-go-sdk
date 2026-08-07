// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateAppConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateAppConfigResponseBody
	GetRequestId() *string
}

type CreateAppConfigResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppConfigResponseBody) SetData(v bool) *CreateAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAppConfigResponseBody) SetRequestId(v string) *CreateAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
