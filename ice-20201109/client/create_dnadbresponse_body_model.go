// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDNADBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInfo(v *CreateDNADBResponseBodyDBInfo) *CreateDNADBResponseBody
	GetDBInfo() *CreateDNADBResponseBodyDBInfo
	SetRequestId(v string) *CreateDNADBResponseBody
	GetRequestId() *string
}

type CreateDNADBResponseBody struct {
	// The details of the media fingerprint library.
	DBInfo *CreateDNADBResponseBodyDBInfo `json:"DBInfo,omitempty" xml:"DBInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDNADBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDNADBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDNADBResponseBody) GetDBInfo() *CreateDNADBResponseBodyDBInfo {
	return s.DBInfo
}

func (s *CreateDNADBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDNADBResponseBody) SetDBInfo(v *CreateDNADBResponseBodyDBInfo) *CreateDNADBResponseBody {
	s.DBInfo = v
	return s
}

func (s *CreateDNADBResponseBody) SetRequestId(v string) *CreateDNADBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDNADBResponseBody) Validate() error {
	if s.DBInfo != nil {
		if err := s.DBInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDNADBResponseBodyDBInfo struct {
	// The ID of the media fingerprint library. We recommend that you save this ID for subsequent calls of other operations.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	DBId *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	// The description of the media fingerprint library.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The model of the media fingerprint library.
	//
	// example:
	//
	// Video
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name of the media fingerprint library.
	//
	// example:
	//
	// example name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the media fingerprint library. After a media fingerprint library is created, it enters the offline state. After the media fingerprint library is processed at the backend, it enters the active state.
	//
	// example:
	//
	// offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDNADBResponseBodyDBInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateDNADBResponseBodyDBInfo) GoString() string {
	return s.String()
}

func (s *CreateDNADBResponseBodyDBInfo) GetDBId() *string {
	return s.DBId
}

func (s *CreateDNADBResponseBodyDBInfo) GetDescription() *string {
	return s.Description
}

func (s *CreateDNADBResponseBodyDBInfo) GetModel() *string {
	return s.Model
}

func (s *CreateDNADBResponseBodyDBInfo) GetName() *string {
	return s.Name
}

func (s *CreateDNADBResponseBodyDBInfo) GetStatus() *string {
	return s.Status
}

func (s *CreateDNADBResponseBodyDBInfo) SetDBId(v string) *CreateDNADBResponseBodyDBInfo {
	s.DBId = &v
	return s
}

func (s *CreateDNADBResponseBodyDBInfo) SetDescription(v string) *CreateDNADBResponseBodyDBInfo {
	s.Description = &v
	return s
}

func (s *CreateDNADBResponseBodyDBInfo) SetModel(v string) *CreateDNADBResponseBodyDBInfo {
	s.Model = &v
	return s
}

func (s *CreateDNADBResponseBodyDBInfo) SetName(v string) *CreateDNADBResponseBodyDBInfo {
	s.Name = &v
	return s
}

func (s *CreateDNADBResponseBodyDBInfo) SetStatus(v string) *CreateDNADBResponseBodyDBInfo {
	s.Status = &v
	return s
}

func (s *CreateDNADBResponseBodyDBInfo) Validate() error {
	return dara.Validate(s)
}
