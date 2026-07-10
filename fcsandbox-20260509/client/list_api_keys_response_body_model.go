// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*ApiKey) *ListApiKeysResponseBody
	GetApiKeys() []*ApiKey
	SetCode(v string) *ListApiKeysResponseBody
	GetCode() *string
	SetMessage(v string) *ListApiKeysResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListApiKeysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApiKeysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApiKeysResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListApiKeysResponseBody
	GetTotal() *int32
}

type ListApiKeysResponseBody struct {
	ApiKeys    []*ApiKey `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	Code       *string   `json:"code,omitempty" xml:"code,omitempty"`
	Message    *string   `json:"message,omitempty" xml:"message,omitempty"`
	PageNumber *int32    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId  *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total      *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBody) GetApiKeys() []*ApiKey {
	return s.ApiKeys
}

func (s *ListApiKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApiKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApiKeysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiKeysResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListApiKeysResponseBody) SetApiKeys(v []*ApiKey) *ListApiKeysResponseBody {
	s.ApiKeys = v
	return s
}

func (s *ListApiKeysResponseBody) SetCode(v string) *ListApiKeysResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiKeysResponseBody) SetMessage(v string) *ListApiKeysResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiKeysResponseBody) SetPageNumber(v int32) *ListApiKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApiKeysResponseBody) SetPageSize(v int32) *ListApiKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApiKeysResponseBody) SetRequestId(v string) *ListApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiKeysResponseBody) SetTotal(v int32) *ListApiKeysResponseBody {
	s.Total = &v
	return s
}

func (s *ListApiKeysResponseBody) Validate() error {
	if s.ApiKeys != nil {
		for _, item := range s.ApiKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
