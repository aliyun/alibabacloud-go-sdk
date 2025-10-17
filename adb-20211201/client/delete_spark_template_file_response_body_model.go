// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSparkTemplateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteSparkTemplateFileResponseBodyData) *DeleteSparkTemplateFileResponseBody
	GetData() *DeleteSparkTemplateFileResponseBodyData
	SetRequestId(v string) *DeleteSparkTemplateFileResponseBody
	GetRequestId() *string
}

type DeleteSparkTemplateFileResponseBody struct {
	// The deletion result.
	Data *DeleteSparkTemplateFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C3A9594F-1D40-4472-A96C-8FB8AA20D38C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSparkTemplateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileResponseBody) GetData() *DeleteSparkTemplateFileResponseBodyData {
	return s.Data
}

func (s *DeleteSparkTemplateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSparkTemplateFileResponseBody) SetData(v *DeleteSparkTemplateFileResponseBodyData) *DeleteSparkTemplateFileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSparkTemplateFileResponseBody) SetRequestId(v string) *DeleteSparkTemplateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSparkTemplateFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSparkTemplateFileResponseBodyData struct {
	// Indicates whether the template file is deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Succeeded *bool `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s DeleteSparkTemplateFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileResponseBodyData) GetSucceeded() *bool {
	return s.Succeeded
}

func (s *DeleteSparkTemplateFileResponseBodyData) SetSucceeded(v bool) *DeleteSparkTemplateFileResponseBodyData {
	s.Succeeded = &v
	return s
}

func (s *DeleteSparkTemplateFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
