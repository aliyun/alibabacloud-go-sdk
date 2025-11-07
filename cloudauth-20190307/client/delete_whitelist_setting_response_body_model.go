// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhitelistSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWhitelistSettingResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteWhitelistSettingResponseBody
	GetResultObject() *bool
}

type DeleteWhitelistSettingResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s DeleteWhitelistSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWhitelistSettingResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteWhitelistSettingResponseBody) SetRequestId(v string) *DeleteWhitelistSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWhitelistSettingResponseBody) SetResultObject(v bool) *DeleteWhitelistSettingResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteWhitelistSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
