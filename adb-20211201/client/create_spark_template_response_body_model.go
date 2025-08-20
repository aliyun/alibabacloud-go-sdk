// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSparkTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSparkTemplateResponseBodyData) *CreateSparkTemplateResponseBody
	GetData() *CreateSparkTemplateResponseBodyData
	SetRequestId(v string) *CreateSparkTemplateResponseBody
	GetRequestId() *string
}

type CreateSparkTemplateResponseBody struct {
	// The creation result.
	Data *CreateSparkTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSparkTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSparkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateResponseBody) GetData() *CreateSparkTemplateResponseBodyData {
	return s.Data
}

func (s *CreateSparkTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSparkTemplateResponseBody) SetData(v *CreateSparkTemplateResponseBodyData) *CreateSparkTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateSparkTemplateResponseBody) SetRequestId(v string) *CreateSparkTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSparkTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSparkTemplateResponseBodyData struct {
	// Indicates whether the application template is created. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// example:
	//
	// True
	Succeeded *bool `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s CreateSparkTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSparkTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateResponseBodyData) GetSucceeded() *bool {
	return s.Succeeded
}

func (s *CreateSparkTemplateResponseBodyData) SetSucceeded(v bool) *CreateSparkTemplateResponseBodyData {
	s.Succeeded = &v
	return s
}

func (s *CreateSparkTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
