// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaskingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]*string) *DeleteMaskingRulesResponseBody
	GetData() map[string]*string
	SetMessage(v string) *DeleteMaskingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMaskingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteMaskingRulesResponseBody
	GetSuccess() *string
}

type DeleteMaskingRulesResponseBody struct {
	Data      map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMaskingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaskingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaskingRulesResponseBody) GetData() map[string]*string {
	return s.Data
}

func (s *DeleteMaskingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMaskingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaskingRulesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteMaskingRulesResponseBody) SetData(v map[string]*string) *DeleteMaskingRulesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMaskingRulesResponseBody) SetMessage(v string) *DeleteMaskingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMaskingRulesResponseBody) SetRequestId(v string) *DeleteMaskingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaskingRulesResponseBody) SetSuccess(v string) *DeleteMaskingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMaskingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
