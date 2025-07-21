// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyVersion(v *CreateKeyVersionResponseBodyKeyVersion) *CreateKeyVersionResponseBody
	GetKeyVersion() *CreateKeyVersionResponseBodyKeyVersion
	SetRequestId(v string) *CreateKeyVersionResponseBody
	GetRequestId() *string
}

type CreateKeyVersionResponseBody struct {
	// The metadata of the version.
	KeyVersion *CreateKeyVersionResponseBodyKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// b96f250a-4b75-498c-91be-22c6928f85be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponseBody) GetKeyVersion() *CreateKeyVersionResponseBodyKeyVersion {
	return s.KeyVersion
}

func (s *CreateKeyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKeyVersionResponseBody) SetKeyVersion(v *CreateKeyVersionResponseBodyKeyVersion) *CreateKeyVersionResponseBody {
	s.KeyVersion = v
	return s
}

func (s *CreateKeyVersionResponseBody) SetRequestId(v string) *CreateKeyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeyVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateKeyVersionResponseBodyKeyVersion struct {
	// The date and time when the version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-02T10:38:27Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the version.
	//
	// example:
	//
	// c0a3d5dc-0b47-4199-a050-b289349a****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s CreateKeyVersionResponseBodyKeyVersion) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyVersionResponseBodyKeyVersion) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponseBodyKeyVersion) GetCreationDate() *string {
	return s.CreationDate
}

func (s *CreateKeyVersionResponseBodyKeyVersion) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateKeyVersionResponseBodyKeyVersion) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetCreationDate(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.CreationDate = &v
	return s
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetKeyId(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.KeyId = &v
	return s
}

func (s *CreateKeyVersionResponseBodyKeyVersion) SetKeyVersionId(v string) *CreateKeyVersionResponseBodyKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *CreateKeyVersionResponseBodyKeyVersion) Validate() error {
	return dara.Validate(s)
}
