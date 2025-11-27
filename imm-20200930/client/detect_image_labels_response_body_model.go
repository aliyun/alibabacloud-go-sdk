// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*Label) *DetectImageLabelsResponseBody
	GetLabels() []*Label
	SetRequestId(v string) *DetectImageLabelsResponseBody
	GetRequestId() *string
}

type DetectImageLabelsResponseBody struct {
	// The list of labels detected.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 91AC8C98-0F36-49D2-8290-742E24DF1F69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageLabelsResponseBody) GetLabels() []*Label {
	return s.Labels
}

func (s *DetectImageLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageLabelsResponseBody) SetLabels(v []*Label) *DetectImageLabelsResponseBody {
	s.Labels = v
	return s
}

func (s *DetectImageLabelsResponseBody) SetRequestId(v string) *DetectImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageLabelsResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
