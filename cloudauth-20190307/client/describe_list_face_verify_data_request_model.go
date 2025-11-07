// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListFaceVerifyDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGmtEnd(v int64) *DescribeListFaceVerifyDataRequest
	GetGmtEnd() *int64
	SetGmtStart(v int64) *DescribeListFaceVerifyDataRequest
	GetGmtStart() *int64
	SetName(v string) *DescribeListFaceVerifyDataRequest
	GetName() *string
	SetSceneId(v int64) *DescribeListFaceVerifyDataRequest
	GetSceneId() *int64
}

type DescribeListFaceVerifyDataRequest struct {
	// End time of the query.
	//
	// example:
	//
	// 1760630399999
	GmtEnd *int64 `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// Start time of the query.
	//
	// example:
	//
	// 1760025600000
	GmtStart *int64 `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	// Product Code, currently deprecated.
	//
	// example:
	//
	// Liveness
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000000339
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeListFaceVerifyDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyDataRequest) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *DescribeListFaceVerifyDataRequest) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *DescribeListFaceVerifyDataRequest) GetName() *string {
	return s.Name
}

func (s *DescribeListFaceVerifyDataRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeListFaceVerifyDataRequest) SetGmtEnd(v int64) *DescribeListFaceVerifyDataRequest {
	s.GmtEnd = &v
	return s
}

func (s *DescribeListFaceVerifyDataRequest) SetGmtStart(v int64) *DescribeListFaceVerifyDataRequest {
	s.GmtStart = &v
	return s
}

func (s *DescribeListFaceVerifyDataRequest) SetName(v string) *DescribeListFaceVerifyDataRequest {
	s.Name = &v
	return s
}

func (s *DescribeListFaceVerifyDataRequest) SetSceneId(v int64) *DescribeListFaceVerifyDataRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeListFaceVerifyDataRequest) Validate() error {
	return dara.Validate(s)
}
