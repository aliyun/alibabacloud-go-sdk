// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeAutoShowListTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *InitializeAutoShowListTaskResponseBody
	GetCasterId() *string
	SetRequestId(v string) *InitializeAutoShowListTaskResponseBody
	GetRequestId() *string
	SetStreamList(v string) *InitializeAutoShowListTaskResponseBody
	GetStreamList() *string
}

type InitializeAutoShowListTaskResponseBody struct {
	// The production studio ID.
	//
	// >This ID can be used as a request parameter for querying production studio stream URLs, starting the production studio, adding video resources to the production studio, adding layouts to the production studio, querying the layout list of the production studio, adding components to the production studio, and adding programs to the production studio.
	//
	// example:
	//
	// b4810848-bcf9-4aef-bd4a-e6bba2d9****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of output addresses of the production studio.
	//
	// - videoFormat: the streaming URL format.
	//
	// - outputStreamUrl: the stream pulling URL.
	//
	// - transcodeConfig: the transcoding resolution description of the stream pulling URL.
	//
	// example:
	//
	// [{"videoFormat":"flv","outputStreamUrl":"http://example.aliyundoc.com","transcodeConfig":"original"}]
	StreamList *string `json:"StreamList,omitempty" xml:"StreamList,omitempty"`
}

func (s InitializeAutoShowListTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeAutoShowListTaskResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeAutoShowListTaskResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *InitializeAutoShowListTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeAutoShowListTaskResponseBody) GetStreamList() *string {
	return s.StreamList
}

func (s *InitializeAutoShowListTaskResponseBody) SetCasterId(v string) *InitializeAutoShowListTaskResponseBody {
	s.CasterId = &v
	return s
}

func (s *InitializeAutoShowListTaskResponseBody) SetRequestId(v string) *InitializeAutoShowListTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeAutoShowListTaskResponseBody) SetStreamList(v string) *InitializeAutoShowListTaskResponseBody {
	s.StreamList = &v
	return s
}

func (s *InitializeAutoShowListTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
