// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceivedFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListReceivedFileRequest
	GetLimit() *int32
	SetMarker(v string) *ListReceivedFileRequest
	GetMarker() *string
}

type ListReceivedFileRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// eym***
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListReceivedFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReceivedFileRequest) GoString() string {
	return s.String()
}

func (s *ListReceivedFileRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListReceivedFileRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListReceivedFileRequest) SetLimit(v int32) *ListReceivedFileRequest {
	s.Limit = &v
	return s
}

func (s *ListReceivedFileRequest) SetMarker(v string) *ListReceivedFileRequest {
	s.Marker = &v
	return s
}

func (s *ListReceivedFileRequest) Validate() error {
	return dara.Validate(s)
}
