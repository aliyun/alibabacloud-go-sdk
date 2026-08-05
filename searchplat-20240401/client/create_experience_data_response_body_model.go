// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperienceDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateExperienceDataResponseBody
	GetRequestId() *string
	SetResult(v *CreateExperienceDataResponseBodyResult) *CreateExperienceDataResponseBody
	GetResult() *CreateExperienceDataResponseBodyResult
}

type CreateExperienceDataResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7ACFD0C5-61E4-5DEA-A995-8279BB99C7E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *CreateExperienceDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateExperienceDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExperienceDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperienceDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExperienceDataResponseBody) GetResult() *CreateExperienceDataResponseBodyResult {
	return s.Result
}

func (s *CreateExperienceDataResponseBody) SetRequestId(v string) *CreateExperienceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperienceDataResponseBody) SetResult(v *CreateExperienceDataResponseBodyResult) *CreateExperienceDataResponseBody {
	s.Result = v
	return s
}

func (s *CreateExperienceDataResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExperienceDataResponseBodyResult struct {
	// The data content type.
	//
	// example:
	//
	// text
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1729665694
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The data size.
	//
	// example:
	//
	// 100
	DataSize *int64 `json:"dataSize,omitempty" xml:"dataSize,omitempty"`
	// The data type.
	//
	// example:
	//
	// file
	DataType *bool `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The data content.
	//
	// example:
	//
	// http://xxx
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// The data ID.
	//
	// example:
	//
	// 1222212
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The data name.
	//
	// example:
	//
	// test
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
	// 1729665694
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateExperienceDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateExperienceDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateExperienceDataResponseBodyResult) GetContentType() *string {
	return s.ContentType
}

func (s *CreateExperienceDataResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *CreateExperienceDataResponseBodyResult) GetDataSize() *int64 {
	return s.DataSize
}

func (s *CreateExperienceDataResponseBodyResult) GetDataType() *bool {
	return s.DataType
}

func (s *CreateExperienceDataResponseBodyResult) GetDataValue() *string {
	return s.DataValue
}

func (s *CreateExperienceDataResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateExperienceDataResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateExperienceDataResponseBodyResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateExperienceDataResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *CreateExperienceDataResponseBodyResult) SetContentType(v string) *CreateExperienceDataResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetCreated(v int64) *CreateExperienceDataResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetDataSize(v int64) *CreateExperienceDataResponseBodyResult {
	s.DataSize = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetDataType(v bool) *CreateExperienceDataResponseBodyResult {
	s.DataType = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetDataValue(v string) *CreateExperienceDataResponseBodyResult {
	s.DataValue = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetId(v int64) *CreateExperienceDataResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetName(v string) *CreateExperienceDataResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetServiceType(v string) *CreateExperienceDataResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) SetUpdated(v int64) *CreateExperienceDataResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateExperienceDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
