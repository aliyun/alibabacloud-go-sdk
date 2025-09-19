// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageVulWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteImageVulWhitelistResponseBody
	GetCode() *string
	SetData(v bool) *DeleteImageVulWhitelistResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteImageVulWhitelistResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteImageVulWhitelistResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteImageVulWhitelistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteImageVulWhitelistResponseBody
	GetSuccess() *bool
}

type DeleteImageVulWhitelistResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the image vulnerability whitelist is deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageVulWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageVulWhitelistResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteImageVulWhitelistResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteImageVulWhitelistResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteImageVulWhitelistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteImageVulWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageVulWhitelistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteImageVulWhitelistResponseBody) SetCode(v string) *DeleteImageVulWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageVulWhitelistResponseBody) SetData(v bool) *DeleteImageVulWhitelistResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteImageVulWhitelistResponseBody) SetHttpStatusCode(v int32) *DeleteImageVulWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteImageVulWhitelistResponseBody) SetMessage(v string) *DeleteImageVulWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageVulWhitelistResponseBody) SetRequestId(v string) *DeleteImageVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageVulWhitelistResponseBody) SetSuccess(v bool) *DeleteImageVulWhitelistResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteImageVulWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
