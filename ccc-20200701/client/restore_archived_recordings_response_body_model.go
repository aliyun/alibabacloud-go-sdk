// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreArchivedRecordingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RestoreArchivedRecordingsResponseBody
	GetCode() *string
	SetData(v []*RestoreArchivedRecordingsResponseBodyData) *RestoreArchivedRecordingsResponseBody
	GetData() []*RestoreArchivedRecordingsResponseBodyData
	SetHttpStatusCode(v int32) *RestoreArchivedRecordingsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RestoreArchivedRecordingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestoreArchivedRecordingsResponseBody
	GetRequestId() *string
}

type RestoreArchivedRecordingsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*RestoreArchivedRecordingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance 0 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F8066648-5D95-55AB-ACD3-2F4AD3BEA715
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreArchivedRecordingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreArchivedRecordingsResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreArchivedRecordingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *RestoreArchivedRecordingsResponseBody) GetData() []*RestoreArchivedRecordingsResponseBodyData {
	return s.Data
}

func (s *RestoreArchivedRecordingsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RestoreArchivedRecordingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestoreArchivedRecordingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreArchivedRecordingsResponseBody) SetCode(v string) *RestoreArchivedRecordingsResponseBody {
	s.Code = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBody) SetData(v []*RestoreArchivedRecordingsResponseBodyData) *RestoreArchivedRecordingsResponseBody {
	s.Data = v
	return s
}

func (s *RestoreArchivedRecordingsResponseBody) SetHttpStatusCode(v int32) *RestoreArchivedRecordingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBody) SetMessage(v string) *RestoreArchivedRecordingsResponseBody {
	s.Message = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBody) SetRequestId(v string) *RestoreArchivedRecordingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type RestoreArchivedRecordingsResponseBodyData struct {
	// example:
	//
	// job-25920271311543****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// False
	Exists *string `json:"Exists,omitempty" xml:"Exists,omitempty"`
	// example:
	//
	// Progressing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Standard
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s RestoreArchivedRecordingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestoreArchivedRecordingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestoreArchivedRecordingsResponseBodyData) GetContactId() *string {
	return s.ContactId
}

func (s *RestoreArchivedRecordingsResponseBodyData) GetExists() *string {
	return s.Exists
}

func (s *RestoreArchivedRecordingsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *RestoreArchivedRecordingsResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *RestoreArchivedRecordingsResponseBodyData) SetContactId(v string) *RestoreArchivedRecordingsResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBodyData) SetExists(v string) *RestoreArchivedRecordingsResponseBodyData {
	s.Exists = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBodyData) SetStatus(v string) *RestoreArchivedRecordingsResponseBodyData {
	s.Status = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBodyData) SetStorageType(v string) *RestoreArchivedRecordingsResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *RestoreArchivedRecordingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
