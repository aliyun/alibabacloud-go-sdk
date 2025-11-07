// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhitelistSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWhitelistSettingResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateWhitelistSettingResponseBody
	GetResultObject() *bool
}

type CreateWhitelistSettingResponseBody struct {
	// The ID of this request.
	//
	// example:
	//
	// CF4979D3-060F-5336-BD9F-0D46ECB66B22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result, whether the creation was successful:
	//
	// - true: Success
	//
	// - false: Failure
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s CreateWhitelistSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWhitelistSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWhitelistSettingResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateWhitelistSettingResponseBody) SetRequestId(v string) *CreateWhitelistSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWhitelistSettingResponseBody) SetResultObject(v bool) *CreateWhitelistSettingResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateWhitelistSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
