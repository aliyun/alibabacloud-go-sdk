// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProduceEditingProjectVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *ProduceEditingProjectVideoResponseBody
	GetMediaId() *string
	SetProjectId(v string) *ProduceEditingProjectVideoResponseBody
	GetProjectId() *string
	SetRequestId(v string) *ProduceEditingProjectVideoResponseBody
	GetRequestId() *string
}

type ProduceEditingProjectVideoResponseBody struct {
	// The ID of the produced video.
	//
	// > 	- This parameter is returned for each request.
	//
	// > 	- If a value is returned for this parameter, the video production task is being asynchronously processed.
	//
	// example:
	//
	// 006204a11bb386bb25491f95f****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// fb2101bf24b4cb318787dc****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProduceEditingProjectVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProduceEditingProjectVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ProduceEditingProjectVideoResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *ProduceEditingProjectVideoResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *ProduceEditingProjectVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProduceEditingProjectVideoResponseBody) SetMediaId(v string) *ProduceEditingProjectVideoResponseBody {
	s.MediaId = &v
	return s
}

func (s *ProduceEditingProjectVideoResponseBody) SetProjectId(v string) *ProduceEditingProjectVideoResponseBody {
	s.ProjectId = &v
	return s
}

func (s *ProduceEditingProjectVideoResponseBody) SetRequestId(v string) *ProduceEditingProjectVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProduceEditingProjectVideoResponseBody) Validate() error {
	return dara.Validate(s)
}
