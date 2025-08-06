// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMezzaninesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistVideoIds(v []*string) *DeleteMezzaninesResponseBody
	GetNonExistVideoIds() []*string
	SetRequestId(v string) *DeleteMezzaninesResponseBody
	GetRequestId() *string
	SetUnRemoveableVideoIds(v []*string) *DeleteMezzaninesResponseBody
	GetUnRemoveableVideoIds() []*string
}

type DeleteMezzaninesResponseBody struct {
	// The IDs of the audio or video files that do not exist.
	NonExistVideoIds []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the audio or video files whose source files cannot be deleted.
	//
	// >  In most cases, source files cannot be deleted if they are used for original-quality playback or you do not have required permissions to delete them. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	UnRemoveableVideoIds []*string `json:"UnRemoveableVideoIds,omitempty" xml:"UnRemoveableVideoIds,omitempty" type:"Repeated"`
}

func (s DeleteMezzaninesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMezzaninesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMezzaninesResponseBody) GetNonExistVideoIds() []*string {
	return s.NonExistVideoIds
}

func (s *DeleteMezzaninesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMezzaninesResponseBody) GetUnRemoveableVideoIds() []*string {
	return s.UnRemoveableVideoIds
}

func (s *DeleteMezzaninesResponseBody) SetNonExistVideoIds(v []*string) *DeleteMezzaninesResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *DeleteMezzaninesResponseBody) SetRequestId(v string) *DeleteMezzaninesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMezzaninesResponseBody) SetUnRemoveableVideoIds(v []*string) *DeleteMezzaninesResponseBody {
	s.UnRemoveableVideoIds = v
	return s
}

func (s *DeleteMezzaninesResponseBody) Validate() error {
	return dara.Validate(s)
}
