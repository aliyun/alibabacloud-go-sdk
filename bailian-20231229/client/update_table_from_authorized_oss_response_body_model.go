// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableFromAuthorizedOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTableFromAuthorizedOssResponseBody
	GetCode() *string
	SetData(v *UpdateTableFromAuthorizedOssResponseBodyData) *UpdateTableFromAuthorizedOssResponseBody
	GetData() *UpdateTableFromAuthorizedOssResponseBodyData
	SetMessage(v string) *UpdateTableFromAuthorizedOssResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTableFromAuthorizedOssResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateTableFromAuthorizedOssResponseBody
	GetStatus() *string
	SetSuccess(v bool) *UpdateTableFromAuthorizedOssResponseBody
	GetSuccess() *bool
}

type UpdateTableFromAuthorizedOssResponseBody struct {
	// Error status code.
	//
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data field of the response.
	Data *UpdateTableFromAuthorizedOssResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request.
	//
	// example:
	//
	// 7BA8ADD9-53D6-53F0-918F-A1E776AD230E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTableFromAuthorizedOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableFromAuthorizedOssResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableFromAuthorizedOssResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTableFromAuthorizedOssResponseBody) GetData() *UpdateTableFromAuthorizedOssResponseBodyData {
	return s.Data
}

func (s *UpdateTableFromAuthorizedOssResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTableFromAuthorizedOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableFromAuthorizedOssResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateTableFromAuthorizedOssResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTableFromAuthorizedOssResponseBody) SetCode(v string) *UpdateTableFromAuthorizedOssResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBody) SetData(v *UpdateTableFromAuthorizedOssResponseBodyData) *UpdateTableFromAuthorizedOssResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBody) SetMessage(v string) *UpdateTableFromAuthorizedOssResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBody) SetRequestId(v string) *UpdateTableFromAuthorizedOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBody) SetStatus(v string) *UpdateTableFromAuthorizedOssResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBody) SetSuccess(v bool) *UpdateTableFromAuthorizedOssResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTableFromAuthorizedOssResponseBodyData struct {
	// Current table status. After upload, the status usually changes to TO_IMPORT. This means the system has accepted the uploaded file and is waiting to schedule it into the data table.
	//
	// example:
	//
	// TO_IMPORT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Same as the input TableId parameter.
	//
	// example:
	//
	// table_df96ebd5da8640e5a0991b3d15f39d4d_12792097
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s UpdateTableFromAuthorizedOssResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableFromAuthorizedOssResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTableFromAuthorizedOssResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateTableFromAuthorizedOssResponseBodyData) GetTableId() *string {
	return s.TableId
}

func (s *UpdateTableFromAuthorizedOssResponseBodyData) SetStatus(v string) *UpdateTableFromAuthorizedOssResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBodyData) SetTableId(v string) *UpdateTableFromAuthorizedOssResponseBodyData {
	s.TableId = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponseBodyData) Validate() error {
	return dara.Validate(s)
}
