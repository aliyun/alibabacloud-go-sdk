// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloneVersion(v int32) *CloneVersionResponseBody
	GetCloneVersion() *int32
	SetOriginVersion(v int32) *CloneVersionResponseBody
	GetOriginVersion() *int32
	SetRequestId(v string) *CloneVersionResponseBody
	GetRequestId() *string
}

type CloneVersionResponseBody struct {
	// The version number of the cloned version.
	//
	// example:
	//
	// 1
	CloneVersion *int32 `json:"CloneVersion,omitempty" xml:"CloneVersion,omitempty"`
	// The version number that was cloned.
	//
	// example:
	//
	// 0
	OriginVersion *int32 `json:"OriginVersion,omitempty" xml:"OriginVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CloneVersionResponseBody) GetCloneVersion() *int32 {
	return s.CloneVersion
}

func (s *CloneVersionResponseBody) GetOriginVersion() *int32 {
	return s.OriginVersion
}

func (s *CloneVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneVersionResponseBody) SetCloneVersion(v int32) *CloneVersionResponseBody {
	s.CloneVersion = &v
	return s
}

func (s *CloneVersionResponseBody) SetOriginVersion(v int32) *CloneVersionResponseBody {
	s.OriginVersion = &v
	return s
}

func (s *CloneVersionResponseBody) SetRequestId(v string) *CloneVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
