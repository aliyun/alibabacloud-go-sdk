// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthPercentageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHealthPercentageResponseBody
	GetCode() *string
	SetData(v []*GetHealthPercentageResponseBodyData) *GetHealthPercentageResponseBody
	GetData() []*GetHealthPercentageResponseBodyData
	SetMessage(v string) *GetHealthPercentageResponseBody
	GetMessage() *string
}

type GetHealthPercentageResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetHealthPercentageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetHealthPercentageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHealthPercentageResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHealthPercentageResponseBody) GetData() []*GetHealthPercentageResponseBodyData {
	return s.Data
}

func (s *GetHealthPercentageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHealthPercentageResponseBody) SetCode(v string) *GetHealthPercentageResponseBody {
	s.Code = &v
	return s
}

func (s *GetHealthPercentageResponseBody) SetData(v []*GetHealthPercentageResponseBodyData) *GetHealthPercentageResponseBody {
	s.Data = v
	return s
}

func (s *GetHealthPercentageResponseBody) SetMessage(v string) *GetHealthPercentageResponseBody {
	s.Message = &v
	return s
}

func (s *GetHealthPercentageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHealthPercentageResponseBodyData struct {
	// example:
	//
	// health
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetHealthPercentageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHealthPercentageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetHealthPercentageResponseBodyData) GetValue() *int64 {
	return s.Value
}

func (s *GetHealthPercentageResponseBodyData) SetType(v string) *GetHealthPercentageResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetHealthPercentageResponseBodyData) SetValue(v int64) *GetHealthPercentageResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetHealthPercentageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
