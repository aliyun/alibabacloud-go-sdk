// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDSRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDSRecordList(v []*QueryDSRecordResponseBodyDSRecordList) *QueryDSRecordResponseBody
	GetDSRecordList() []*QueryDSRecordResponseBodyDSRecordList
	SetRequestId(v string) *QueryDSRecordResponseBody
	GetRequestId() *string
}

type QueryDSRecordResponseBody struct {
	DSRecordList []*QueryDSRecordResponseBodyDSRecordList `json:"DSRecordList,omitempty" xml:"DSRecordList,omitempty" type:"Repeated"`
	// example:
	//
	// 814B2AF0-ED6F-4C13-B41C-8AC0B1023583
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDSRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDSRecordResponseBody) GetDSRecordList() []*QueryDSRecordResponseBodyDSRecordList {
	return s.DSRecordList
}

func (s *QueryDSRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDSRecordResponseBody) SetDSRecordList(v []*QueryDSRecordResponseBodyDSRecordList) *QueryDSRecordResponseBody {
	s.DSRecordList = v
	return s
}

func (s *QueryDSRecordResponseBody) SetRequestId(v string) *QueryDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDSRecordResponseBody) Validate() error {
	if s.DSRecordList != nil {
		for _, item := range s.DSRecordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDSRecordResponseBodyDSRecordList struct {
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// f58fa917424383934c7b0cf1a90f61d692745680fa06f5ecdbe0924e86de9598
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// example:
	//
	// 2
	DigestType *int32 `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	// example:
	//
	// 1
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
}

func (s QueryDSRecordResponseBodyDSRecordList) String() string {
	return dara.Prettify(s)
}

func (s QueryDSRecordResponseBodyDSRecordList) GoString() string {
	return s.String()
}

func (s *QueryDSRecordResponseBodyDSRecordList) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *QueryDSRecordResponseBodyDSRecordList) GetDigest() *string {
	return s.Digest
}

func (s *QueryDSRecordResponseBodyDSRecordList) GetDigestType() *int32 {
	return s.DigestType
}

func (s *QueryDSRecordResponseBodyDSRecordList) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetAlgorithm(v int32) *QueryDSRecordResponseBodyDSRecordList {
	s.Algorithm = &v
	return s
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetDigest(v string) *QueryDSRecordResponseBodyDSRecordList {
	s.Digest = &v
	return s
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetDigestType(v int32) *QueryDSRecordResponseBodyDSRecordList {
	s.DigestType = &v
	return s
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetKeyTag(v int32) *QueryDSRecordResponseBodyDSRecordList {
	s.KeyTag = &v
	return s
}

func (s *QueryDSRecordResponseBodyDSRecordList) Validate() error {
	return dara.Validate(s)
}
