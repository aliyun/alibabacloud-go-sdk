// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSparkTemplateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateSparkTemplateFileResponseBodyData) *UpdateSparkTemplateFileResponseBody
	GetData() *UpdateSparkTemplateFileResponseBodyData
	SetRequestId(v string) *UpdateSparkTemplateFileResponseBody
	GetRequestId() *string
}

type UpdateSparkTemplateFileResponseBody struct {
	// The update result.
	Data *UpdateSparkTemplateFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C3A9594F-1D40-4472-A96C-8FB8AA20D38C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSparkTemplateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSparkTemplateFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileResponseBody) GetData() *UpdateSparkTemplateFileResponseBodyData {
	return s.Data
}

func (s *UpdateSparkTemplateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSparkTemplateFileResponseBody) SetData(v *UpdateSparkTemplateFileResponseBodyData) *UpdateSparkTemplateFileResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSparkTemplateFileResponseBody) SetRequestId(v string) *UpdateSparkTemplateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSparkTemplateFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateSparkTemplateFileResponseBodyData struct {
	// Indicates whether the application template is updated.
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

func (s UpdateSparkTemplateFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSparkTemplateFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileResponseBodyData) GetSucceeded() *bool {
	return s.Succeeded
}

func (s *UpdateSparkTemplateFileResponseBodyData) SetSucceeded(v bool) *UpdateSparkTemplateFileResponseBodyData {
	s.Succeeded = &v
	return s
}

func (s *UpdateSparkTemplateFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
