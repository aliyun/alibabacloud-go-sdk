// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQualityRuleResponseBody
	GetCode() *string
	SetData(v string) *DeleteQualityRuleResponseBody
	GetData() *string
	SetMessage(v string) *DeleteQualityRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityRuleResponseBody
	GetSuccess() *bool
}

type DeleteQualityRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQualityRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteQualityRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityRuleResponseBody) SetCode(v string) *DeleteQualityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetData(v string) *DeleteQualityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetMessage(v string) *DeleteQualityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetRequestId(v string) *DeleteQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetSuccess(v bool) *DeleteQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
