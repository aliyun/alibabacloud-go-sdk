// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsumerAPIKeyResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConsumerAPIKeyResult
	GetCode() *string
	SetData(v *ConsumerAPIKey) *ConsumerAPIKeyResult
	GetData() *ConsumerAPIKey
	SetRequestId(v string) *ConsumerAPIKeyResult
	GetRequestId() *string
}

type ConsumerAPIKeyResult struct {
	// `SUCCESS` indicates that the request was successful. If the request fails, this field returns an error type, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Details of the consumer API key.
	//
	// example:
	//
	// {}
	Data *ConsumerAPIKey `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request ID for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ConsumerAPIKeyResult) String() string {
	return dara.Prettify(s)
}

func (s ConsumerAPIKeyResult) GoString() string {
	return s.String()
}

func (s *ConsumerAPIKeyResult) GetCode() *string {
	return s.Code
}

func (s *ConsumerAPIKeyResult) GetData() *ConsumerAPIKey {
	return s.Data
}

func (s *ConsumerAPIKeyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsumerAPIKeyResult) SetCode(v string) *ConsumerAPIKeyResult {
	s.Code = &v
	return s
}

func (s *ConsumerAPIKeyResult) SetData(v *ConsumerAPIKey) *ConsumerAPIKeyResult {
	s.Data = v
	return s
}

func (s *ConsumerAPIKeyResult) SetRequestId(v string) *ConsumerAPIKeyResult {
	s.RequestId = &v
	return s
}

func (s *ConsumerAPIKeyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
