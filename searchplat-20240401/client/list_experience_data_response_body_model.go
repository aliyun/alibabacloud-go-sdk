// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperienceDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListExperienceDataResponseBody
	GetRequestId() *string
	SetResult(v []*ListExperienceDataResponseBodyResult) *ListExperienceDataResponseBody
	GetResult() []*ListExperienceDataResponseBodyResult
}

type ListExperienceDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FDSS_1232
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*ListExperienceDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListExperienceDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExperienceDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperienceDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExperienceDataResponseBody) GetResult() []*ListExperienceDataResponseBodyResult {
	return s.Result
}

func (s *ListExperienceDataResponseBody) SetRequestId(v string) *ListExperienceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperienceDataResponseBody) SetResult(v []*ListExperienceDataResponseBodyResult) *ListExperienceDataResponseBody {
	s.Result = v
	return s
}

func (s *ListExperienceDataResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExperienceDataResponseBodyResult struct {
	// **The content type.**.
	//
	// example:
	//
	// pdf
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// **The creation time.**.
	//
	// example:
	//
	// 12313123123
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// **The data size.**.
	//
	// example:
	//
	// 100
	DataSize *int64 `json:"dataSize,omitempty" xml:"dataSize,omitempty"`
	// The data type. Valid values:
	//
	// - file
	//
	// - url.
	//
	// example:
	//
	// file
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The data value.
	//
	// example:
	//
	// oss://bucket/xxx.pdf
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// ID
	//
	// example:
	//
	// 9bd21be8
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The name.
	//
	// example:
	//
	// xxx.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service type.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1232131231
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListExperienceDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListExperienceDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListExperienceDataResponseBodyResult) GetContentType() *string {
	return s.ContentType
}

func (s *ListExperienceDataResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *ListExperienceDataResponseBodyResult) GetDataSize() *int64 {
	return s.DataSize
}

func (s *ListExperienceDataResponseBodyResult) GetDataType() *string {
	return s.DataType
}

func (s *ListExperienceDataResponseBodyResult) GetDataValue() *string {
	return s.DataValue
}

func (s *ListExperienceDataResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListExperienceDataResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListExperienceDataResponseBodyResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListExperienceDataResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *ListExperienceDataResponseBodyResult) SetContentType(v string) *ListExperienceDataResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetCreated(v int64) *ListExperienceDataResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetDataSize(v int64) *ListExperienceDataResponseBodyResult {
	s.DataSize = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetDataType(v string) *ListExperienceDataResponseBodyResult {
	s.DataType = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetDataValue(v string) *ListExperienceDataResponseBodyResult {
	s.DataValue = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetId(v int64) *ListExperienceDataResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetName(v string) *ListExperienceDataResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetServiceType(v string) *ListExperienceDataResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) SetUpdated(v int64) *ListExperienceDataResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListExperienceDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
