// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateBackupResponseBodyData) *CreateBackupResponseBody
	GetData() *CreateBackupResponseBodyData
	SetMessage(v string) *CreateBackupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBackupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBackupResponseBody
	GetSuccess() *bool
}

type CreateBackupResponseBody struct {
	Data *CreateBackupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) GetData() *CreateBackupResponseBodyData {
	return s.Data
}

func (s *CreateBackupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBackupResponseBody) SetData(v *CreateBackupResponseBodyData) *CreateBackupResponseBody {
	s.Data = v
	return s
}

func (s *CreateBackupResponseBody) SetMessage(v string) *CreateBackupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackupResponseBody) SetRequestId(v string) *CreateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponseBody) SetSuccess(v bool) *CreateBackupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBackupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateBackupResponseBodyData struct {
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
}

func (s CreateBackupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBodyData) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateBackupResponseBodyData) SetBackupSetId(v string) *CreateBackupResponseBodyData {
	s.BackupSetId = &v
	return s
}

func (s *CreateBackupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
