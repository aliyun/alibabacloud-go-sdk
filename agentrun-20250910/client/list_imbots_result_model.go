// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIMBotsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIMBotsResult
	GetCode() *string
	SetData(v *ListIMBotsOutput) *ListIMBotsResult
	GetData() *ListIMBotsOutput
	SetRequestId(v string) *ListIMBotsResult
	GetRequestId() *string
}

type ListIMBotsResult struct {
	Code      *string           `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListIMBotsOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string           `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListIMBotsResult) String() string {
	return dara.Prettify(s)
}

func (s ListIMBotsResult) GoString() string {
	return s.String()
}

func (s *ListIMBotsResult) GetCode() *string {
	return s.Code
}

func (s *ListIMBotsResult) GetData() *ListIMBotsOutput {
	return s.Data
}

func (s *ListIMBotsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIMBotsResult) SetCode(v string) *ListIMBotsResult {
	s.Code = &v
	return s
}

func (s *ListIMBotsResult) SetData(v *ListIMBotsOutput) *ListIMBotsResult {
	s.Data = v
	return s
}

func (s *ListIMBotsResult) SetRequestId(v string) *ListIMBotsResult {
	s.RequestId = &v
	return s
}

func (s *ListIMBotsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
