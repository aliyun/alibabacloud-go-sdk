// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFileResponseBody
	GetCode() *string
	SetData(v *DeleteFileResponseBodyData) *DeleteFileResponseBody
	GetData() *DeleteFileResponseBodyData
	SetMessage(v string) *DeleteFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFileResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteFileResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteFileResponseBody
	GetSuccess() *bool
}

type DeleteFileResponseBody struct {
	// example:
	//
	// InvalidParameter
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Current file status does not support delete.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFileResponseBody) GetData() *DeleteFileResponseBodyData {
	return s.Data
}

func (s *DeleteFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFileResponseBody) SetCode(v string) *DeleteFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFileResponseBody) SetData(v *DeleteFileResponseBodyData) *DeleteFileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteFileResponseBody) SetMessage(v string) *DeleteFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetStatus(v string) *DeleteFileResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteFileResponseBodyData struct {
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s DeleteFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *DeleteFileResponseBodyData) SetFileId(v string) *DeleteFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *DeleteFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
