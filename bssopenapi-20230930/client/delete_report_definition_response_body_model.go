// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReportDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteReportDefinitionResponseBody
	GetData() *bool
	SetMetadata(v interface{}) *DeleteReportDefinitionResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *DeleteReportDefinitionResponseBody
	GetRequestId() *string
}

type DeleteReportDefinitionResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReportDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteReportDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReportDefinitionResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteReportDefinitionResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *DeleteReportDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteReportDefinitionResponseBody) SetData(v bool) *DeleteReportDefinitionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteReportDefinitionResponseBody) SetMetadata(v interface{}) *DeleteReportDefinitionResponseBody {
	s.Metadata = v
	return s
}

func (s *DeleteReportDefinitionResponseBody) SetRequestId(v string) *DeleteReportDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteReportDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
