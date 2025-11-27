// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteADSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteADSettingResponseBody
	GetRequestId() *string
}

type DeleteADSettingResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteADSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteADSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteADSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteADSettingResponseBody) SetRequestId(v string) *DeleteADSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteADSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
