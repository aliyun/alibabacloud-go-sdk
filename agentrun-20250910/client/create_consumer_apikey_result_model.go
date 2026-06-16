// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAPIKeyResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConsumerAPIKeyResult
	GetCode() *string
	SetData(v *CreateConsumerAPIKeyOutput) *CreateConsumerAPIKeyResult
	GetData() *CreateConsumerAPIKeyOutput
	SetRequestId(v string) *CreateConsumerAPIKeyResult
	GetRequestId() *string
}

type CreateConsumerAPIKeyResult struct {
	// `SUCCESS` for a successful operation; otherwise, the corresponding error type.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Details of the new consumer API key, including the complete key.
	//
	// example:
	//
	// {}
	Data *CreateConsumerAPIKeyOutput `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request id for issue tracking.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateConsumerAPIKeyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAPIKeyResult) GoString() string {
	return s.String()
}

func (s *CreateConsumerAPIKeyResult) GetCode() *string {
	return s.Code
}

func (s *CreateConsumerAPIKeyResult) GetData() *CreateConsumerAPIKeyOutput {
	return s.Data
}

func (s *CreateConsumerAPIKeyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerAPIKeyResult) SetCode(v string) *CreateConsumerAPIKeyResult {
	s.Code = &v
	return s
}

func (s *CreateConsumerAPIKeyResult) SetData(v *CreateConsumerAPIKeyOutput) *CreateConsumerAPIKeyResult {
	s.Data = v
	return s
}

func (s *CreateConsumerAPIKeyResult) SetRequestId(v string) *CreateConsumerAPIKeyResult {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerAPIKeyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
