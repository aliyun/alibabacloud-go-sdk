// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDriveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportDriveResponseBody
	GetCode() *string
	SetDriveId(v string) *ImportDriveResponseBody
	GetDriveId() *string
	SetMessage(v string) *ImportDriveResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportDriveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportDriveResponseBody
	GetSuccess() *bool
}

type ImportDriveResponseBody struct {
	// The modification result. A value of `success` indicates success. Otherwise, an error message is returned.
	//
	// example:
	//
	// Response code. 200 indicates success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The newly created or reused target drive ID. A success response only indicates that the import task has been accepted.
	//
	// example:
	//
	// dri-aaaa****
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// The response message.
	//
	// example:
	//
	// Response message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A87DBB05-653A-5E4B-B72B-5F4A1E07****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportDriveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportDriveResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDriveResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportDriveResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *ImportDriveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportDriveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportDriveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportDriveResponseBody) SetCode(v string) *ImportDriveResponseBody {
	s.Code = &v
	return s
}

func (s *ImportDriveResponseBody) SetDriveId(v string) *ImportDriveResponseBody {
	s.DriveId = &v
	return s
}

func (s *ImportDriveResponseBody) SetMessage(v string) *ImportDriveResponseBody {
	s.Message = &v
	return s
}

func (s *ImportDriveResponseBody) SetRequestId(v string) *ImportDriveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportDriveResponseBody) SetSuccess(v bool) *ImportDriveResponseBody {
	s.Success = &v
	return s
}

func (s *ImportDriveResponseBody) Validate() error {
	return dara.Validate(s)
}
