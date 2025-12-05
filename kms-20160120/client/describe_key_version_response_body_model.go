// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyVersion(v *DescribeKeyVersionResponseBodyKeyVersion) *DescribeKeyVersionResponseBody
	GetKeyVersion() *DescribeKeyVersionResponseBodyKeyVersion
	SetRequestId(v string) *DescribeKeyVersionResponseBody
	GetRequestId() *string
}

type DescribeKeyVersionResponseBody struct {
	// The metadata of the CMK version.
	KeyVersion *DescribeKeyVersionResponseBodyKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKeyVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponseBody) GetKeyVersion() *DescribeKeyVersionResponseBodyKeyVersion {
	return s.KeyVersion
}

func (s *DescribeKeyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKeyVersionResponseBody) SetKeyVersion(v *DescribeKeyVersionResponseBodyKeyVersion) *DescribeKeyVersionResponseBody {
	s.KeyVersion = v
	return s
}

func (s *DescribeKeyVersionResponseBody) SetRequestId(v string) *DescribeKeyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyVersionResponseBody) Validate() error {
	if s.KeyVersion != nil {
		if err := s.KeyVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKeyVersionResponseBodyKeyVersion struct {
	// The date and time when the CMK version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-03-25T10:42:40Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter in the request to an alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s DescribeKeyVersionResponseBodyKeyVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyVersionResponseBodyKeyVersion) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) GetCreationDate() *string {
	return s.CreationDate
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetCreationDate(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.CreationDate = &v
	return s
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetKeyId(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) SetKeyVersionId(v string) *DescribeKeyVersionResponseBodyKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *DescribeKeyVersionResponseBodyKeyVersion) Validate() error {
	return dara.Validate(s)
}
