// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ImportDataSourcesResponseBodyData) *ImportDataSourcesResponseBody
	GetData() *ImportDataSourcesResponseBodyData
	SetRequestId(v string) *ImportDataSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportDataSourcesResponseBody
	GetSuccess() *bool
}

type ImportDataSourcesResponseBody struct {
	// The information about the imported data sources.
	Data *ImportDataSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDataSourcesResponseBody) GetData() *ImportDataSourcesResponseBodyData {
	return s.Data
}

func (s *ImportDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportDataSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportDataSourcesResponseBody) SetData(v *ImportDataSourcesResponseBodyData) *ImportDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ImportDataSourcesResponseBody) SetRequestId(v string) *ImportDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportDataSourcesResponseBody) SetSuccess(v bool) *ImportDataSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ImportDataSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImportDataSourcesResponseBodyData struct {
	// The reason why the data sources failed to be imported. If the data sources were imported, this parameter is left empty.
	//
	// example:
	//
	// Data source DEV XXX already exists
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the data sources were imported. Valid values:
	//
	// 	- true: All data sources were imported.
	//
	// 	- false: Specific data sources failed to be imported. You can troubleshoot issues based on the Message parameter.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ImportDataSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportDataSourcesResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ImportDataSourcesResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *ImportDataSourcesResponseBodyData) SetMessage(v string) *ImportDataSourcesResponseBodyData {
	s.Message = &v
	return s
}

func (s *ImportDataSourcesResponseBodyData) SetStatus(v bool) *ImportDataSourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ImportDataSourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
