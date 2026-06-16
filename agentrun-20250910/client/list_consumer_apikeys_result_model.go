// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerAPIKeysResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerAPIKeysResult
	GetCode() *string
	SetData(v *ListConsumerAPIKeysOutput) *ListConsumerAPIKeysResult
	GetData() *ListConsumerAPIKeysOutput
	SetRequestId(v string) *ListConsumerAPIKeysResult
	GetRequestId() *string
}

type ListConsumerAPIKeysResult struct {
	// The request status. `SUCCESS` indicates that the request was successful. Otherwise, this field returns the error type.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of consumer API keys.
	//
	// example:
	//
	// {}
	Data *ListConsumerAPIKeysOutput `json:"data,omitempty" xml:"data,omitempty"`
	// The unique request ID. Use this ID for issue tracking.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConsumerAPIKeysResult) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerAPIKeysResult) GoString() string {
	return s.String()
}

func (s *ListConsumerAPIKeysResult) GetCode() *string {
	return s.Code
}

func (s *ListConsumerAPIKeysResult) GetData() *ListConsumerAPIKeysOutput {
	return s.Data
}

func (s *ListConsumerAPIKeysResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerAPIKeysResult) SetCode(v string) *ListConsumerAPIKeysResult {
	s.Code = &v
	return s
}

func (s *ListConsumerAPIKeysResult) SetData(v *ListConsumerAPIKeysOutput) *ListConsumerAPIKeysResult {
	s.Data = v
	return s
}

func (s *ListConsumerAPIKeysResult) SetRequestId(v string) *ListConsumerAPIKeysResult {
	s.RequestId = &v
	return s
}

func (s *ListConsumerAPIKeysResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
