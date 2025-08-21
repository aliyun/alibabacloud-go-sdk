// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryTagsRequest
	GetAppKey() *int64
	SetClientKey(v string) *QueryTagsRequest
	GetClientKey() *string
	SetKeyType(v string) *QueryTagsRequest
	GetKeyType() *string
}

type QueryTagsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e2ba19de97604f55b165576****
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
}

func (s QueryTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTagsRequest) GoString() string {
	return s.String()
}

func (s *QueryTagsRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryTagsRequest) GetClientKey() *string {
	return s.ClientKey
}

func (s *QueryTagsRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *QueryTagsRequest) SetAppKey(v int64) *QueryTagsRequest {
	s.AppKey = &v
	return s
}

func (s *QueryTagsRequest) SetClientKey(v string) *QueryTagsRequest {
	s.ClientKey = &v
	return s
}

func (s *QueryTagsRequest) SetKeyType(v string) *QueryTagsRequest {
	s.KeyType = &v
	return s
}

func (s *QueryTagsRequest) Validate() error {
	return dara.Validate(s)
}
