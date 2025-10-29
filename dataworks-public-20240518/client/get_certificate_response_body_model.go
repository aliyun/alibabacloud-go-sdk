// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v *GetCertificateResponseBodyCertificate) *GetCertificateResponseBody
	GetCertificate() *GetCertificateResponseBodyCertificate
	SetRequestId(v string) *GetCertificateResponseBody
	GetRequestId() *string
}

type GetCertificateResponseBody struct {
	// The details of the certificate file.
	Certificate *GetCertificateResponseBodyCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBody) GetCertificate() *GetCertificateResponseBodyCertificate {
	return s.Certificate
}

func (s *GetCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificateResponseBody) SetCertificate(v *GetCertificateResponseBodyCertificate) *GetCertificateResponseBody {
	s.Certificate = v
	return s
}

func (s *GetCertificateResponseBody) SetRequestId(v string) *GetCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateResponseBody) Validate() error {
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCertificateResponseBodyCertificate struct {
	// The time when the certificate file was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1730217600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the certificate file.
	//
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a file
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The size of the certificate file, in bytes.
	//
	// example:
	//
	// 77549
	FileSizeInBytes *int64 `json:"FileSizeInBytes,omitempty" xml:"FileSizeInBytes,omitempty"`
	// The ID of the certificate file.
	//
	// example:
	//
	// 676303114031776
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the certificate file.
	//
	// example:
	//
	// ca1.crt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the workspace to which the certificate file belongs.
	//
	// example:
	//
	// 177161
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetCertificateResponseBodyCertificate) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateResponseBodyCertificate) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBodyCertificate) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCertificateResponseBodyCertificate) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetCertificateResponseBodyCertificate) GetDescription() *string {
	return s.Description
}

func (s *GetCertificateResponseBodyCertificate) GetFileSizeInBytes() *int64 {
	return s.FileSizeInBytes
}

func (s *GetCertificateResponseBodyCertificate) GetId() *int64 {
	return s.Id
}

func (s *GetCertificateResponseBodyCertificate) GetName() *string {
	return s.Name
}

func (s *GetCertificateResponseBodyCertificate) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCertificateResponseBodyCertificate) SetCreateTime(v int64) *GetCertificateResponseBodyCertificate {
	s.CreateTime = &v
	return s
}

func (s *GetCertificateResponseBodyCertificate) SetCreateUser(v string) *GetCertificateResponseBodyCertificate {
	s.CreateUser = &v
	return s
}

func (s *GetCertificateResponseBodyCertificate) SetDescription(v string) *GetCertificateResponseBodyCertificate {
	s.Description = &v
	return s
}

func (s *GetCertificateResponseBodyCertificate) SetFileSizeInBytes(v int64) *GetCertificateResponseBodyCertificate {
	s.FileSizeInBytes = &v
	return s
}

func (s *GetCertificateResponseBodyCertificate) SetId(v int64) *GetCertificateResponseBodyCertificate {
	s.Id = &v
	return s
}

func (s *GetCertificateResponseBodyCertificate) SetName(v string) *GetCertificateResponseBodyCertificate {
	s.Name = &v
	return s
}

func (s *GetCertificateResponseBodyCertificate) SetProjectId(v int64) *GetCertificateResponseBodyCertificate {
	s.ProjectId = &v
	return s
}

func (s *GetCertificateResponseBodyCertificate) Validate() error {
	return dara.Validate(s)
}
