// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRestartApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *BatchRestartApplicationsRequest
	GetAppIds() *string
	SetNamespaceId(v string) *BatchRestartApplicationsRequest
	GetNamespaceId() *string
}

type BatchRestartApplicationsRequest struct {
	// The application IDs. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 00893b3a-0e24-45ed-be0c-1f20e07b****
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s BatchRestartApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRestartApplicationsRequest) GoString() string {
	return s.String()
}

func (s *BatchRestartApplicationsRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *BatchRestartApplicationsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *BatchRestartApplicationsRequest) SetAppIds(v string) *BatchRestartApplicationsRequest {
	s.AppIds = &v
	return s
}

func (s *BatchRestartApplicationsRequest) SetNamespaceId(v string) *BatchRestartApplicationsRequest {
	s.NamespaceId = &v
	return s
}

func (s *BatchRestartApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
