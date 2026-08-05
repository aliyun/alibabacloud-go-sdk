// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperienceDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetExperienceDataResponseBody
	GetRequestId() *string
	SetResult(v *GetExperienceDataResponseBodyResult) *GetExperienceDataResponseBody
	GetResult() *GetExperienceDataResponseBodyResult
}

type GetExperienceDataResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5950143C-B8F0-5758-A08A-66F302FD587F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *GetExperienceDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetExperienceDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperienceDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperienceDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperienceDataResponseBody) GetResult() *GetExperienceDataResponseBodyResult {
	return s.Result
}

func (s *GetExperienceDataResponseBody) SetRequestId(v string) *GetExperienceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperienceDataResponseBody) SetResult(v *GetExperienceDataResponseBodyResult) *GetExperienceDataResponseBody {
	s.Result = v
	return s
}

func (s *GetExperienceDataResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExperienceDataResponseBodyResult struct {
	// The file type.
	//
	// - pdf
	//
	// - text
	//
	// - html
	//
	// - doc.
	//
	// example:
	//
	// text
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1745806839720
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The data size.
	//
	// example:
	//
	// 100
	DataSize *int64 `json:"dataSize,omitempty" xml:"dataSize,omitempty"`
	// The data type.
	//
	// - file
	//
	// - url.
	//
	// example:
	//
	// file
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The data content.
	//
	// - When dataType is set to file, this field is the OSS address of the file.
	//
	// - When dataType is set to url, this field is the HTTP URL of the data.
	//
	// example:
	//
	// http://xxx
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// The data ID.
	//
	// example:
	//
	// 1877
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The data name. This parameter is required when dataType is set to file.
	//
	// example:
	//
	// a.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service type.
	//
	// - document-analyze.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1729684154
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s GetExperienceDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetExperienceDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetExperienceDataResponseBodyResult) GetContentType() *string {
	return s.ContentType
}

func (s *GetExperienceDataResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *GetExperienceDataResponseBodyResult) GetDataSize() *int64 {
	return s.DataSize
}

func (s *GetExperienceDataResponseBodyResult) GetDataType() *string {
	return s.DataType
}

func (s *GetExperienceDataResponseBodyResult) GetDataValue() *string {
	return s.DataValue
}

func (s *GetExperienceDataResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetExperienceDataResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetExperienceDataResponseBodyResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetExperienceDataResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *GetExperienceDataResponseBodyResult) SetContentType(v string) *GetExperienceDataResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetCreated(v int64) *GetExperienceDataResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetDataSize(v int64) *GetExperienceDataResponseBodyResult {
	s.DataSize = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetDataType(v string) *GetExperienceDataResponseBodyResult {
	s.DataType = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetDataValue(v string) *GetExperienceDataResponseBodyResult {
	s.DataValue = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetId(v int64) *GetExperienceDataResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetName(v string) *GetExperienceDataResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetServiceType(v string) *GetExperienceDataResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) SetUpdated(v int64) *GetExperienceDataResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetExperienceDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
