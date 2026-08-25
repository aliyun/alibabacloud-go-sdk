// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVolumesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVolumesResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListVolumesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListVolumesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListVolumesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVolumesResponseBody
	GetRequestId() *string
	SetVolumes(v []*E2BVolume) *ListVolumesResponseBody
	GetVolumes() []*E2BVolume
}

type ListVolumesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token used to retrieve more results. You do not need to specify this parameter for the first request. For subsequent requests, use the token returned in the previous response.
	//
	// example:
	//
	// qxGrXje86XMrYQ51aJMy
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of volumes.
	Volumes []*E2BVolume `json:"volumes,omitempty" xml:"volumes,omitempty" type:"Repeated"`
}

func (s ListVolumesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVolumesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVolumesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVolumesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVolumesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVolumesResponseBody) GetVolumes() []*E2BVolume {
	return s.Volumes
}

func (s *ListVolumesResponseBody) SetCode(v string) *ListVolumesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVolumesResponseBody) SetMaxResults(v int32) *ListVolumesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVolumesResponseBody) SetMessage(v string) *ListVolumesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVolumesResponseBody) SetNextToken(v string) *ListVolumesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVolumesResponseBody) SetRequestId(v string) *ListVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVolumesResponseBody) SetVolumes(v []*E2BVolume) *ListVolumesResponseBody {
	s.Volumes = v
	return s
}

func (s *ListVolumesResponseBody) Validate() error {
	if s.Volumes != nil {
		for _, item := range s.Volumes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
