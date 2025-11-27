// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageFacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFaces(v []*Figure) *DetectImageFacesResponseBody
	GetFaces() []*Figure
	SetRequestId(v string) *DetectImageFacesResponseBody
	GetRequestId() *string
}

type DetectImageFacesResponseBody struct {
	// The faces.
	Faces []*Figure `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageFacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBody) GetFaces() []*Figure {
	return s.Faces
}

func (s *DetectImageFacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageFacesResponseBody) SetFaces(v []*Figure) *DetectImageFacesResponseBody {
	s.Faces = v
	return s
}

func (s *DetectImageFacesResponseBody) SetRequestId(v string) *DetectImageFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageFacesResponseBody) Validate() error {
	if s.Faces != nil {
		for _, item := range s.Faces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
