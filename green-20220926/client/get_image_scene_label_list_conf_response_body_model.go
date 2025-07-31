// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageSceneLabelListConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []interface{}) *GetImageSceneLabelListConfResponseBody
	GetData() []interface{}
	SetRequestId(v string) *GetImageSceneLabelListConfResponseBody
	GetRequestId() *string
}

type GetImageSceneLabelListConfResponseBody struct {
	Data []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageSceneLabelListConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageSceneLabelListConfResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelListConfResponseBody) GetData() []interface{} {
	return s.Data
}

func (s *GetImageSceneLabelListConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageSceneLabelListConfResponseBody) SetData(v []interface{}) *GetImageSceneLabelListConfResponseBody {
	s.Data = v
	return s
}

func (s *GetImageSceneLabelListConfResponseBody) SetRequestId(v string) *GetImageSceneLabelListConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageSceneLabelListConfResponseBody) Validate() error {
	return dara.Validate(s)
}
