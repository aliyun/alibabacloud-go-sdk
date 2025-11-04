// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNADBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBList(v []*ListDNADBResponseBodyDBList) *ListDNADBResponseBody
	GetDBList() []*ListDNADBResponseBodyDBList
	SetRequestId(v string) *ListDNADBResponseBody
	GetRequestId() *string
}

type ListDNADBResponseBody struct {
	// The queried media fingerprint libraries.
	DBList []*ListDNADBResponseBodyDBList `json:"DBList,omitempty" xml:"DBList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDNADBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDNADBResponseBody) GoString() string {
	return s.String()
}

func (s *ListDNADBResponseBody) GetDBList() []*ListDNADBResponseBodyDBList {
	return s.DBList
}

func (s *ListDNADBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDNADBResponseBody) SetDBList(v []*ListDNADBResponseBodyDBList) *ListDNADBResponseBody {
	s.DBList = v
	return s
}

func (s *ListDNADBResponseBody) SetRequestId(v string) *ListDNADBResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDNADBResponseBody) Validate() error {
	if s.DBList != nil {
		for _, item := range s.DBList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDNADBResponseBodyDBList struct {
	// The ID of the media fingerprint library.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	DBId *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	// The description of the media fingerprint library.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The model of the media fingerprint library. Valid values:
	//
	// 	- **Video**
	//
	// 	- **Audio**
	//
	// 	- **Image**
	//
	// 	- **Text*	- (supported only in the China (Shanghai) region)
	//
	// example:
	//
	// Video
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name of the media fingerprint library.
	//
	// example:
	//
	// example-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the media fingerprint library. Default value: **offline**. ****Valid values:
	//
	// 	- **offline**: The media fingerprint library is offline.
	//
	// 	- **active**: The media fingerprint library is online.
	//
	// 	- **deleted**: The media fingerprint library is deleted.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDNADBResponseBodyDBList) String() string {
	return dara.Prettify(s)
}

func (s ListDNADBResponseBodyDBList) GoString() string {
	return s.String()
}

func (s *ListDNADBResponseBodyDBList) GetDBId() *string {
	return s.DBId
}

func (s *ListDNADBResponseBodyDBList) GetDescription() *string {
	return s.Description
}

func (s *ListDNADBResponseBodyDBList) GetModel() *string {
	return s.Model
}

func (s *ListDNADBResponseBodyDBList) GetName() *string {
	return s.Name
}

func (s *ListDNADBResponseBodyDBList) GetStatus() *string {
	return s.Status
}

func (s *ListDNADBResponseBodyDBList) SetDBId(v string) *ListDNADBResponseBodyDBList {
	s.DBId = &v
	return s
}

func (s *ListDNADBResponseBodyDBList) SetDescription(v string) *ListDNADBResponseBodyDBList {
	s.Description = &v
	return s
}

func (s *ListDNADBResponseBodyDBList) SetModel(v string) *ListDNADBResponseBodyDBList {
	s.Model = &v
	return s
}

func (s *ListDNADBResponseBodyDBList) SetName(v string) *ListDNADBResponseBodyDBList {
	s.Name = &v
	return s
}

func (s *ListDNADBResponseBodyDBList) SetStatus(v string) *ListDNADBResponseBodyDBList {
	s.Status = &v
	return s
}

func (s *ListDNADBResponseBodyDBList) Validate() error {
	return dara.Validate(s)
}
