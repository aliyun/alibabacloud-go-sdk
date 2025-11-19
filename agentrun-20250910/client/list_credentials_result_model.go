// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCredentialsResult
	GetCode() *string
	SetData(v *ListCredentialsOutput) *ListCredentialsResult
	GetData() *ListCredentialsOutput
	SetRequestId(v string) *ListCredentialsResult
	GetRequestId() *string
}

type ListCredentialsResult struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListCredentialsOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListCredentialsResult) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResult) GoString() string {
	return s.String()
}

func (s *ListCredentialsResult) GetCode() *string {
	return s.Code
}

func (s *ListCredentialsResult) GetData() *ListCredentialsOutput {
	return s.Data
}

func (s *ListCredentialsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCredentialsResult) SetCode(v string) *ListCredentialsResult {
	s.Code = &v
	return s
}

func (s *ListCredentialsResult) SetData(v *ListCredentialsOutput) *ListCredentialsResult {
	s.Data = v
	return s
}

func (s *ListCredentialsResult) SetRequestId(v string) *ListCredentialsResult {
	s.RequestId = &v
	return s
}

func (s *ListCredentialsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
