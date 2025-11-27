// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBodiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBodies(v []*Body) *DetectImageBodiesResponseBody
	GetBodies() []*Body
	SetRequestId(v string) *DetectImageBodiesResponseBody
	GetRequestId() *string
}

type DetectImageBodiesResponseBody struct {
	// The human bodies.
	Bodies []*Body `json:"Bodies,omitempty" xml:"Bodies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 501339F9-4B70-0CE2-AB8C-866C********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageBodiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBodiesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBody) GetBodies() []*Body {
	return s.Bodies
}

func (s *DetectImageBodiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageBodiesResponseBody) SetBodies(v []*Body) *DetectImageBodiesResponseBody {
	s.Bodies = v
	return s
}

func (s *DetectImageBodiesResponseBody) SetRequestId(v string) *DetectImageBodiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageBodiesResponseBody) Validate() error {
	if s.Bodies != nil {
		for _, item := range s.Bodies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
