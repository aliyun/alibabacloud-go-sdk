// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistTemplateLinkedInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeWhitelistTemplateLinkedInstanceResponseBody
	GetCode() *string
	SetData(v *DescribeWhitelistTemplateLinkedInstanceResponseBodyData) *DescribeWhitelistTemplateLinkedInstanceResponseBody
	GetData() *DescribeWhitelistTemplateLinkedInstanceResponseBodyData
	SetHttpStatusCode(v int32) *DescribeWhitelistTemplateLinkedInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeWhitelistTemplateLinkedInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeWhitelistTemplateLinkedInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeWhitelistTemplateLinkedInstanceResponseBody
	GetSuccess() *bool
}

type DescribeWhitelistTemplateLinkedInstanceResponseBody struct {
	// The response code returned. Valid values:
	//
	// 	- **200**: success
	//
	// 	- **400**: client error
	//
	// 	- **401**: identity authentication failed
	//
	// 	- **404**: request page not found
	//
	// 	- **500**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeWhitelistTemplateLinkedInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned. Valid values:
	//
	// 	- **200**: success
	//
	// 	- **400**: client error
	//
	// 	- **500**: server error
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9F8C06AD-3F37-57A0-ABBF-ABD7824F55CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
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

func (s DescribeWhitelistTemplateLinkedInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateLinkedInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) GetData() *DescribeWhitelistTemplateLinkedInstanceResponseBodyData {
	return s.Data
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) SetCode(v string) *DescribeWhitelistTemplateLinkedInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) SetData(v *DescribeWhitelistTemplateLinkedInstanceResponseBodyData) *DescribeWhitelistTemplateLinkedInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) SetHttpStatusCode(v int32) *DescribeWhitelistTemplateLinkedInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) SetMessage(v string) *DescribeWhitelistTemplateLinkedInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) SetRequestId(v string) *DescribeWhitelistTemplateLinkedInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) SetSuccess(v bool) *DescribeWhitelistTemplateLinkedInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWhitelistTemplateLinkedInstanceResponseBodyData struct {
	// The information about the instance.
	InsName []*string `json:"InsName,omitempty" xml:"InsName,omitempty" type:"Repeated"`
	// The ID of the whitelist template.
	//
	// example:
	//
	// 412
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeWhitelistTemplateLinkedInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateLinkedInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBodyData) GetInsName() []*string {
	return s.InsName
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBodyData) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBodyData) SetInsName(v []*string) *DescribeWhitelistTemplateLinkedInstanceResponseBodyData {
	s.InsName = v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBodyData) SetTemplateId(v int32) *DescribeWhitelistTemplateLinkedInstanceResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
