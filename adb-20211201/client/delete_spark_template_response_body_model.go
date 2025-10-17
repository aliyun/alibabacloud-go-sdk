// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSparkTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteSparkTemplateResponseBodyData) *DeleteSparkTemplateResponseBody
	GetData() *DeleteSparkTemplateResponseBodyData
	SetRequestId(v string) *DeleteSparkTemplateResponseBody
	GetRequestId() *string
}

type DeleteSparkTemplateResponseBody struct {
	// The returned result.
	Data *DeleteSparkTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSparkTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateResponseBody) GetData() *DeleteSparkTemplateResponseBodyData {
	return s.Data
}

func (s *DeleteSparkTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSparkTemplateResponseBody) SetData(v *DeleteSparkTemplateResponseBodyData) *DeleteSparkTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSparkTemplateResponseBody) SetRequestId(v string) *DeleteSparkTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSparkTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSparkTemplateResponseBodyData struct {
	// Indicates whether the request was successful. Valid values:
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

func (s DeleteSparkTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateResponseBodyData) GetSucceeded() *bool {
	return s.Succeeded
}

func (s *DeleteSparkTemplateResponseBodyData) SetSucceeded(v bool) *DeleteSparkTemplateResponseBodyData {
	s.Succeeded = &v
	return s
}

func (s *DeleteSparkTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
