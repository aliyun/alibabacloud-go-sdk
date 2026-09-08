// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingEncryptionAlgorithmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithms(v []*ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) *ListDataMaskingEncryptionAlgorithmsResponseBody
	GetAlgorithms() []*ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms
	SetRequestId(v string) *ListDataMaskingEncryptionAlgorithmsResponseBody
	GetRequestId() *string
}

type ListDataMaskingEncryptionAlgorithmsResponseBody struct {
	Algorithms []*ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms `json:"Algorithms,omitempty" xml:"Algorithms,omitempty" type:"Repeated"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataMaskingEncryptionAlgorithmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingEncryptionAlgorithmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBody) GetAlgorithms() []*ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms {
	return s.Algorithms
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBody) SetAlgorithms(v []*ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) *ListDataMaskingEncryptionAlgorithmsResponseBody {
	s.Algorithms = v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBody) SetRequestId(v string) *ListDataMaskingEncryptionAlgorithmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBody) Validate() error {
	if s.Algorithms != nil {
		for _, item := range s.Algorithms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms struct {
	// example:
	//
	// data_masking_not_running
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 实例未处于运行状态
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// AES_256_GCM
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) GoString() string {
	return s.String()
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) GetName() *string {
	return s.Name
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) SetErrorCode(v string) *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms {
	s.ErrorCode = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) SetErrorMessage(v string) *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) SetName(v string) *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms {
	s.Name = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponseBodyAlgorithms) Validate() error {
	return dara.Validate(s)
}
