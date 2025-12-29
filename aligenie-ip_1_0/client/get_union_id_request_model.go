// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnionIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodeKey(v string) *GetUnionIdRequest
	GetEncodeKey() *string
	SetEncodeType(v string) *GetUnionIdRequest
	GetEncodeType() *string
	SetId(v string) *GetUnionIdRequest
	GetId() *string
	SetIdType(v string) *GetUnionIdRequest
	GetIdType() *string
}

type GetUnionIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 125****0946
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62a319****abdc
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEVICE_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s GetUnionIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetUnionIdRequest) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetUnionIdRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetUnionIdRequest) GetId() *string {
	return s.Id
}

func (s *GetUnionIdRequest) GetIdType() *string {
	return s.IdType
}

func (s *GetUnionIdRequest) SetEncodeKey(v string) *GetUnionIdRequest {
	s.EncodeKey = &v
	return s
}

func (s *GetUnionIdRequest) SetEncodeType(v string) *GetUnionIdRequest {
	s.EncodeType = &v
	return s
}

func (s *GetUnionIdRequest) SetId(v string) *GetUnionIdRequest {
	s.Id = &v
	return s
}

func (s *GetUnionIdRequest) SetIdType(v string) *GetUnionIdRequest {
	s.IdType = &v
	return s
}

func (s *GetUnionIdRequest) Validate() error {
	return dara.Validate(s)
}
