// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDetectPornDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectPornData(v *DescribeLiveDetectPornDataResponseBodyDetectPornData) *DescribeLiveDetectPornDataResponseBody
	GetDetectPornData() *DescribeLiveDetectPornDataResponseBodyDetectPornData
	SetRequestId(v string) *DescribeLiveDetectPornDataResponseBody
	GetRequestId() *string
}

type DescribeLiveDetectPornDataResponseBody struct {
	DetectPornData *DescribeLiveDetectPornDataResponseBodyDetectPornData `json:"DetectPornData,omitempty" xml:"DetectPornData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDetectPornDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectPornDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataResponseBody) GetDetectPornData() *DescribeLiveDetectPornDataResponseBodyDetectPornData {
	return s.DetectPornData
}

func (s *DescribeLiveDetectPornDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDetectPornDataResponseBody) SetDetectPornData(v *DescribeLiveDetectPornDataResponseBodyDetectPornData) *DescribeLiveDetectPornDataResponseBody {
	s.DetectPornData = v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBody) SetRequestId(v string) *DescribeLiveDetectPornDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBody) Validate() error {
	if s.DetectPornData != nil {
		if err := s.DetectPornData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDetectPornDataResponseBodyDetectPornData struct {
	DataModule []*DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDetectPornDataResponseBodyDetectPornData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectPornDataResponseBodyDetectPornData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornData) GetDataModule() []*DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	return s.DataModule
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornData) SetDataModule(v []*DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) *DescribeLiveDetectPornDataResponseBodyDetectPornData {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornData) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule struct {
	App       *string `json:"App,omitempty" xml:"App,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Fee       *string `json:"Fee,omitempty" xml:"Fee,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Scene     *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Stream    *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetApp() *string {
	return s.App
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetCount() *int64 {
	return s.Count
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetFee() *string {
	return s.Fee
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetScene() *string {
	return s.Scene
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetStream() *string {
	return s.Stream
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetApp(v string) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.App = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetCount(v int64) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.Count = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetDomain(v string) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetFee(v string) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.Fee = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetRegion(v string) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.Region = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetScene(v string) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.Scene = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetStream(v string) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.Stream = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) SetTimeStamp(v string) *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseBodyDetectPornDataDataModule) Validate() error {
	return dara.Validate(s)
}
