// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type SegmentHDSkyRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentHDSkyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyRequest) SetImageURL(v string) *SegmentHDSkyRequest {
	s.ImageURL = &v
	return s
}

type SegmentHDSkyResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentHDSkyResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentHDSkyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponse) SetRequestId(v string) *SegmentHDSkyResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentHDSkyResponse) SetData(v *SegmentHDSkyResponseData) *SegmentHDSkyResponse {
	s.Data = v
	return s
}

type SegmentHDSkyResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentHDSkyResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyResponseData) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponseData) SetImageURL(v string) *SegmentHDSkyResponseData {
	s.ImageURL = &v
	return s
}

type SegmentHDSkyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentHDSkyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHDSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentHDCommonImageRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
}

func (s SegmentHDCommonImageRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageRequest) SetImageUrl(v string) *SegmentHDCommonImageRequest {
	s.ImageUrl = &v
	return s
}

type SegmentHDCommonImageResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentHDCommonImageResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentHDCommonImageResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponse) SetRequestId(v string) *SegmentHDCommonImageResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentHDCommonImageResponse) SetData(v *SegmentHDCommonImageResponseData) *SegmentHDCommonImageResponse {
	s.Data = v
	return s
}

type SegmentHDCommonImageResponseData struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
}

func (s SegmentHDCommonImageResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageResponseData) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponseData) SetImageUrl(v string) *SegmentHDCommonImageResponseData {
	s.ImageUrl = &v
	return s
}

type SegmentHDCommonImageAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
}

func (s SegmentHDCommonImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDCommonImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageAdvanceRequest) SetImageUrlObject(v io.Reader) *SegmentHDCommonImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type SegmentSkinRequest struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s SegmentSkinRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkinRequest) SetURL(v string) *SegmentSkinRequest {
	s.URL = &v
	return s
}

type SegmentSkinResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentSkinResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentSkinResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinResponse) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponse) SetRequestId(v string) *SegmentSkinResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentSkinResponse) SetData(v *SegmentSkinResponseData) *SegmentSkinResponse {
	s.Data = v
	return s
}

type SegmentSkinResponseData struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s SegmentSkinResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinResponseData) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponseData) SetURL(v string) *SegmentSkinResponseData {
	s.URL = &v
	return s
}

type SegmentSkinAdvanceRequest struct {
	URLObject io.Reader `json:"URLObject,omitempty" xml:"URLObject,omitempty" require:"true"`
}

func (s SegmentSkinAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkinAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkinAdvanceRequest) SetURLObject(v io.Reader) *SegmentSkinAdvanceRequest {
	s.URLObject = v
	return s
}

type ChangeSkyRequest struct {
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	ReplaceImageURL *string `json:"ReplaceImageURL,omitempty" xml:"ReplaceImageURL,omitempty" require:"true"`
}

func (s ChangeSkyRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyRequest) GoString() string {
	return s.String()
}

func (s *ChangeSkyRequest) SetImageURL(v string) *ChangeSkyRequest {
	s.ImageURL = &v
	return s
}

func (s *ChangeSkyRequest) SetReplaceImageURL(v string) *ChangeSkyRequest {
	s.ReplaceImageURL = &v
	return s
}

type ChangeSkyResponse struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ChangeSkyResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ChangeSkyResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyResponse) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponse) SetRequestId(v string) *ChangeSkyResponse {
	s.RequestId = &v
	return s
}

func (s *ChangeSkyResponse) SetData(v *ChangeSkyResponseData) *ChangeSkyResponse {
	s.Data = v
	return s
}

type ChangeSkyResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s ChangeSkyResponseData) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyResponseData) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponseData) SetImageURL(v string) *ChangeSkyResponseData {
	s.ImageURL = &v
	return s
}

type ChangeSkyAdvanceRequest struct {
	ImageURLObject  io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ReplaceImageURL *string   `json:"ReplaceImageURL,omitempty" xml:"ReplaceImageURL,omitempty" require:"true"`
}

func (s ChangeSkyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeSkyAdvanceRequest) SetImageURLObject(v io.Reader) *ChangeSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ChangeSkyAdvanceRequest) SetReplaceImageURL(v string) *ChangeSkyAdvanceRequest {
	s.ReplaceImageURL = &v
	return s
}

type SegmentLogoRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentLogoRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoRequest) GoString() string {
	return s.String()
}

func (s *SegmentLogoRequest) SetImageURL(v string) *SegmentLogoRequest {
	s.ImageURL = &v
	return s
}

type SegmentLogoResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentLogoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentLogoResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoResponse) GoString() string {
	return s.String()
}

func (s *SegmentLogoResponse) SetRequestId(v string) *SegmentLogoResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentLogoResponse) SetData(v *SegmentLogoResponseData) *SegmentLogoResponse {
	s.Data = v
	return s
}

type SegmentLogoResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentLogoResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoResponseData) GoString() string {
	return s.String()
}

func (s *SegmentLogoResponseData) SetImageURL(v string) *SegmentLogoResponseData {
	s.ImageURL = &v
	return s
}

type SegmentLogoAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentLogoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentLogoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentLogoAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentLogoAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentSceneRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneRequest) GoString() string {
	return s.String()
}

func (s *SegmentSceneRequest) SetImageURL(v string) *SegmentSceneRequest {
	s.ImageURL = &v
	return s
}

type SegmentSceneResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentSceneResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneResponse) GoString() string {
	return s.String()
}

func (s *SegmentSceneResponse) SetRequestId(v string) *SegmentSceneResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentSceneResponse) SetData(v *SegmentSceneResponseData) *SegmentSceneResponse {
	s.Data = v
	return s
}

type SegmentSceneResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentSceneResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneResponseData) GoString() string {
	return s.String()
}

func (s *SegmentSceneResponseData) SetImageURL(v string) *SegmentSceneResponseData {
	s.ImageURL = &v
	return s
}

type SegmentSceneAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentSceneAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSceneAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSceneAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentSceneAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentFoodRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentFoodRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodRequest) GoString() string {
	return s.String()
}

func (s *SegmentFoodRequest) SetImageURL(v string) *SegmentFoodRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentFoodRequest) SetReturnForm(v string) *SegmentFoodRequest {
	s.ReturnForm = &v
	return s
}

type SegmentFoodResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentFoodResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentFoodResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodResponse) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponse) SetRequestId(v string) *SegmentFoodResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentFoodResponse) SetData(v *SegmentFoodResponseData) *SegmentFoodResponse {
	s.Data = v
	return s
}

type SegmentFoodResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentFoodResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodResponseData) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponseData) SetImageURL(v string) *SegmentFoodResponseData {
	s.ImageURL = &v
	return s
}

type SegmentFoodAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentFoodAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFoodAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFoodAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentFoodAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentFoodAdvanceRequest) SetReturnForm(v string) *SegmentFoodAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentClothRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentClothRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothRequest) SetImageURL(v string) *SegmentClothRequest {
	s.ImageURL = &v
	return s
}

type SegmentClothResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentClothResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentClothResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponse) GoString() string {
	return s.String()
}

func (s *SegmentClothResponse) SetRequestId(v string) *SegmentClothResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentClothResponse) SetData(v *SegmentClothResponseData) *SegmentClothResponse {
	s.Data = v
	return s
}

type SegmentClothResponseData struct {
	Elements []*SegmentClothResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s SegmentClothResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponseData) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseData) SetElements(v []*SegmentClothResponseDataElements) *SegmentClothResponseData {
	s.Elements = v
	return s
}

type SegmentClothResponseDataElements struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentClothResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothResponseDataElements) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseDataElements) SetImageURL(v string) *SegmentClothResponseDataElements {
	s.ImageURL = &v
	return s
}

type SegmentClothAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentClothAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentClothAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentClothAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentAnimalRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentAnimalRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalRequest) GoString() string {
	return s.String()
}

func (s *SegmentAnimalRequest) SetImageURL(v string) *SegmentAnimalRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentAnimalRequest) SetReturnForm(v string) *SegmentAnimalRequest {
	s.ReturnForm = &v
	return s
}

type SegmentAnimalResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentAnimalResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentAnimalResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalResponse) GoString() string {
	return s.String()
}

func (s *SegmentAnimalResponse) SetRequestId(v string) *SegmentAnimalResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentAnimalResponse) SetData(v *SegmentAnimalResponseData) *SegmentAnimalResponse {
	s.Data = v
	return s
}

type SegmentAnimalResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentAnimalResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalResponseData) GoString() string {
	return s.String()
}

func (s *SegmentAnimalResponseData) SetImageURL(v string) *SegmentAnimalResponseData {
	s.ImageURL = &v
	return s
}

type SegmentAnimalAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentAnimalAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentAnimalAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentAnimalAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentAnimalAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentAnimalAdvanceRequest) SetReturnForm(v string) *SegmentAnimalAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentHDBodyRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentHDBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyRequest) SetImageURL(v string) *SegmentHDBodyRequest {
	s.ImageURL = &v
	return s
}

type SegmentHDBodyResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentHDBodyResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentHDBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponse) SetRequestId(v string) *SegmentHDBodyResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentHDBodyResponse) SetData(v *SegmentHDBodyResponseData) *SegmentHDBodyResponse {
	s.Data = v
	return s
}

type SegmentHDBodyResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentHDBodyResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyResponseData) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponseData) SetImageURL(v string) *SegmentHDBodyResponseData {
	s.ImageURL = &v
	return s
}

type SegmentHDBodyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentHDBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHDBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHDBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentSkyRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentSkyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkyRequest) SetImageURL(v string) *SegmentSkyRequest {
	s.ImageURL = &v
	return s
}

type SegmentSkyResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentSkyResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentSkyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyResponse) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponse) SetRequestId(v string) *SegmentSkyResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentSkyResponse) SetData(v *SegmentSkyResponseData) *SegmentSkyResponse {
	s.Data = v
	return s
}

type SegmentSkyResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentSkyResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyResponseData) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponseData) SetImageURL(v string) *SegmentSkyResponseData {
	s.ImageURL = &v
	return s
}

type SegmentSkyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentSkyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetAsyncJobResultResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetRequestId(v string) *GetAsyncJobResultResponse {
	s.RequestId = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetData(v *GetAsyncJobResultResponseData) *GetAsyncJobResultResponse {
	s.Data = v
	return s
}

type GetAsyncJobResultResponseData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty" require:"true"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseData) SetErrorCode(v string) *GetAsyncJobResultResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetErrorMessage(v string) *GetAsyncJobResultResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetJobId(v string) *GetAsyncJobResultResponseData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetResult(v string) *GetAsyncJobResultResponseData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetStatus(v string) *GetAsyncJobResultResponseData {
	s.Status = &v
	return s
}

type SegmentFurnitureRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentFurnitureRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureRequest) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureRequest) SetImageURL(v string) *SegmentFurnitureRequest {
	s.ImageURL = &v
	return s
}

type SegmentFurnitureResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentFurnitureResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentFurnitureResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureResponse) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureResponse) SetRequestId(v string) *SegmentFurnitureResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentFurnitureResponse) SetData(v *SegmentFurnitureResponseData) *SegmentFurnitureResponse {
	s.Data = v
	return s
}

type SegmentFurnitureResponseData struct {
	Elements []*SegmentFurnitureResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s SegmentFurnitureResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureResponseData) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureResponseData) SetElements(v []*SegmentFurnitureResponseDataElements) *SegmentFurnitureResponseData {
	s.Elements = v
	return s
}

type SegmentFurnitureResponseDataElements struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentFurnitureResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureResponseDataElements) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureResponseDataElements) SetImageURL(v string) *SegmentFurnitureResponseDataElements {
	s.ImageURL = &v
	return s
}

type SegmentFurnitureAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentFurnitureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFurnitureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFurnitureAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentFurnitureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RefineMaskRequest struct {
	MaskImageURL *string `json:"MaskImageURL,omitempty" xml:"MaskImageURL,omitempty" require:"true"`
	ImageURL     *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s RefineMaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskRequest) GoString() string {
	return s.String()
}

func (s *RefineMaskRequest) SetMaskImageURL(v string) *RefineMaskRequest {
	s.MaskImageURL = &v
	return s
}

func (s *RefineMaskRequest) SetImageURL(v string) *RefineMaskRequest {
	s.ImageURL = &v
	return s
}

type RefineMaskResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *RefineMaskResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s RefineMaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskResponse) GoString() string {
	return s.String()
}

func (s *RefineMaskResponse) SetRequestId(v string) *RefineMaskResponse {
	s.RequestId = &v
	return s
}

func (s *RefineMaskResponse) SetData(v *RefineMaskResponseData) *RefineMaskResponse {
	s.Data = v
	return s
}

type RefineMaskResponseData struct {
	Elements []*RefineMaskResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s RefineMaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskResponseData) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseData) SetElements(v []*RefineMaskResponseDataElements) *RefineMaskResponseData {
	s.Elements = v
	return s
}

type RefineMaskResponseDataElements struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s RefineMaskResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskResponseDataElements) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseDataElements) SetImageURL(v string) *RefineMaskResponseDataElements {
	s.ImageURL = &v
	return s
}

type RefineMaskAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	MaskImageURL   *string   `json:"MaskImageURL,omitempty" xml:"MaskImageURL,omitempty" require:"true"`
}

func (s RefineMaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefineMaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RefineMaskAdvanceRequest) SetImageURLObject(v io.Reader) *RefineMaskAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RefineMaskAdvanceRequest) SetMaskImageURL(v string) *RefineMaskAdvanceRequest {
	s.MaskImageURL = &v
	return s
}

type ParseFaceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s ParseFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceRequest) GoString() string {
	return s.String()
}

func (s *ParseFaceRequest) SetImageURL(v string) *ParseFaceRequest {
	s.ImageURL = &v
	return s
}

type ParseFaceResponse struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ParseFaceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ParseFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceResponse) GoString() string {
	return s.String()
}

func (s *ParseFaceResponse) SetRequestId(v string) *ParseFaceResponse {
	s.RequestId = &v
	return s
}

func (s *ParseFaceResponse) SetData(v *ParseFaceResponseData) *ParseFaceResponse {
	s.Data = v
	return s
}

type ParseFaceResponseData struct {
	OriginImageURL *string                          `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty" require:"true"`
	Elements       []*ParseFaceResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s ParseFaceResponseData) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceResponseData) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseData) SetOriginImageURL(v string) *ParseFaceResponseData {
	s.OriginImageURL = &v
	return s
}

func (s *ParseFaceResponseData) SetElements(v []*ParseFaceResponseDataElements) *ParseFaceResponseData {
	s.Elements = v
	return s
}

type ParseFaceResponseDataElements struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s ParseFaceResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceResponseDataElements) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseDataElements) SetName(v string) *ParseFaceResponseDataElements {
	s.Name = &v
	return s
}

func (s *ParseFaceResponseDataElements) SetImageURL(v string) *ParseFaceResponseDataElements {
	s.ImageURL = &v
	return s
}

type ParseFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s ParseFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ParseFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ParseFaceAdvanceRequest) SetImageURLObject(v io.Reader) *ParseFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentVehicleRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentVehicleRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleRequest) GoString() string {
	return s.String()
}

func (s *SegmentVehicleRequest) SetImageURL(v string) *SegmentVehicleRequest {
	s.ImageURL = &v
	return s
}

type SegmentVehicleResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentVehicleResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentVehicleResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleResponse) GoString() string {
	return s.String()
}

func (s *SegmentVehicleResponse) SetRequestId(v string) *SegmentVehicleResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentVehicleResponse) SetData(v *SegmentVehicleResponseData) *SegmentVehicleResponse {
	s.Data = v
	return s
}

type SegmentVehicleResponseData struct {
	Elements []*SegmentVehicleResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s SegmentVehicleResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleResponseData) GoString() string {
	return s.String()
}

func (s *SegmentVehicleResponseData) SetElements(v []*SegmentVehicleResponseDataElements) *SegmentVehicleResponseData {
	s.Elements = v
	return s
}

type SegmentVehicleResponseDataElements struct {
	OriginImageURL *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty" require:"true"`
	ImageURL       *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentVehicleResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleResponseDataElements) GoString() string {
	return s.String()
}

func (s *SegmentVehicleResponseDataElements) SetOriginImageURL(v string) *SegmentVehicleResponseDataElements {
	s.OriginImageURL = &v
	return s
}

func (s *SegmentVehicleResponseDataElements) SetImageURL(v string) *SegmentVehicleResponseDataElements {
	s.ImageURL = &v
	return s
}

type SegmentVehicleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentVehicleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentVehicleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentVehicleAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentVehicleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentHairRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentHairRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairRequest) GoString() string {
	return s.String()
}

func (s *SegmentHairRequest) SetImageURL(v string) *SegmentHairRequest {
	s.ImageURL = &v
	return s
}

type SegmentHairResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentHairResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentHairResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairResponse) GoString() string {
	return s.String()
}

func (s *SegmentHairResponse) SetRequestId(v string) *SegmentHairResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentHairResponse) SetData(v *SegmentHairResponseData) *SegmentHairResponse {
	s.Data = v
	return s
}

type SegmentHairResponseData struct {
	Elements []*SegmentHairResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s SegmentHairResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairResponseData) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseData) SetElements(v []*SegmentHairResponseDataElements) *SegmentHairResponseData {
	s.Elements = v
	return s
}

type SegmentHairResponseDataElements struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	X        *int    `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Y        *int    `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	Width    *int    `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height   *int    `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
}

func (s SegmentHairResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairResponseDataElements) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseDataElements) SetImageURL(v string) *SegmentHairResponseDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentHairResponseDataElements) SetX(v int) *SegmentHairResponseDataElements {
	s.X = &v
	return s
}

func (s *SegmentHairResponseDataElements) SetY(v int) *SegmentHairResponseDataElements {
	s.Y = &v
	return s
}

func (s *SegmentHairResponseDataElements) SetWidth(v int) *SegmentHairResponseDataElements {
	s.Width = &v
	return s
}

func (s *SegmentHairResponseDataElements) SetHeight(v int) *SegmentHairResponseDataElements {
	s.Height = &v
	return s
}

type SegmentHairAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentHairAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHairAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHairAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHairAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentFaceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFaceRequest) SetImageURL(v string) *SegmentFaceRequest {
	s.ImageURL = &v
	return s
}

type SegmentFaceResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentFaceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceResponse) GoString() string {
	return s.String()
}

func (s *SegmentFaceResponse) SetRequestId(v string) *SegmentFaceResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentFaceResponse) SetData(v *SegmentFaceResponseData) *SegmentFaceResponse {
	s.Data = v
	return s
}

type SegmentFaceResponseData struct {
	Elements []*SegmentFaceResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s SegmentFaceResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceResponseData) GoString() string {
	return s.String()
}

func (s *SegmentFaceResponseData) SetElements(v []*SegmentFaceResponseDataElements) *SegmentFaceResponseData {
	s.Elements = v
	return s
}

type SegmentFaceResponseDataElements struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	X        *int    `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Y        *int    `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	Width    *int    `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height   *int    `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
}

func (s SegmentFaceResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceResponseDataElements) GoString() string {
	return s.String()
}

func (s *SegmentFaceResponseDataElements) SetImageURL(v string) *SegmentFaceResponseDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentFaceResponseDataElements) SetX(v int) *SegmentFaceResponseDataElements {
	s.X = &v
	return s
}

func (s *SegmentFaceResponseDataElements) SetY(v int) *SegmentFaceResponseDataElements {
	s.Y = &v
	return s
}

func (s *SegmentFaceResponseDataElements) SetWidth(v int) *SegmentFaceResponseDataElements {
	s.Width = &v
	return s
}

func (s *SegmentFaceResponseDataElements) SetHeight(v int) *SegmentFaceResponseDataElements {
	s.Height = &v
	return s
}

type SegmentFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s SegmentFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFaceAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type SegmentHeadRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentHeadRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadRequest) GoString() string {
	return s.String()
}

func (s *SegmentHeadRequest) SetImageURL(v string) *SegmentHeadRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentHeadRequest) SetReturnForm(v string) *SegmentHeadRequest {
	s.ReturnForm = &v
	return s
}

type SegmentHeadResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentHeadResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentHeadResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadResponse) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponse) SetRequestId(v string) *SegmentHeadResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentHeadResponse) SetData(v *SegmentHeadResponseData) *SegmentHeadResponse {
	s.Data = v
	return s
}

type SegmentHeadResponseData struct {
	Elements []*SegmentHeadResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s SegmentHeadResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadResponseData) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseData) SetElements(v []*SegmentHeadResponseDataElements) *SegmentHeadResponseData {
	s.Elements = v
	return s
}

type SegmentHeadResponseDataElements struct {
	X        *int    `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Y        *int    `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	Height   *int    `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	Width    *int    `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
}

func (s SegmentHeadResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadResponseDataElements) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseDataElements) SetX(v int) *SegmentHeadResponseDataElements {
	s.X = &v
	return s
}

func (s *SegmentHeadResponseDataElements) SetY(v int) *SegmentHeadResponseDataElements {
	s.Y = &v
	return s
}

func (s *SegmentHeadResponseDataElements) SetImageURL(v string) *SegmentHeadResponseDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentHeadResponseDataElements) SetHeight(v int) *SegmentHeadResponseDataElements {
	s.Height = &v
	return s
}

func (s *SegmentHeadResponseDataElements) SetWidth(v int) *SegmentHeadResponseDataElements {
	s.Width = &v
	return s
}

type SegmentHeadAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentHeadAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHeadAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHeadAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHeadAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentHeadAdvanceRequest) SetReturnForm(v string) *SegmentHeadAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommodityRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommodityRequest) SetImageURL(v string) *SegmentCommodityRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommodityRequest) SetReturnForm(v string) *SegmentCommodityRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommodityResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentCommodityResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponse) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponse) SetRequestId(v string) *SegmentCommodityResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentCommodityResponse) SetData(v *SegmentCommodityResponseData) *SegmentCommodityResponse {
	s.Data = v
	return s
}

type SegmentCommodityResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentCommodityResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponseData) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponseData) SetImageURL(v string) *SegmentCommodityResponseData {
	s.ImageURL = &v
	return s
}

type SegmentCommodityAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommodityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommodityAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentCommodityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentCommodityAdvanceRequest) SetReturnForm(v string) *SegmentCommodityAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentBodyRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	Async      *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentBodyRequest) SetImageURL(v string) *SegmentBodyRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentBodyRequest) SetAsync(v bool) *SegmentBodyRequest {
	s.Async = &v
	return s
}

func (s *SegmentBodyRequest) SetReturnForm(v string) *SegmentBodyRequest {
	s.ReturnForm = &v
	return s
}

type SegmentBodyResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentBodyResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponse) SetRequestId(v string) *SegmentBodyResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentBodyResponse) SetData(v *SegmentBodyResponseData) *SegmentBodyResponse {
	s.Data = v
	return s
}

type SegmentBodyResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentBodyResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponseData) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponseData) SetImageURL(v string) *SegmentBodyResponseData {
	s.ImageURL = &v
	return s
}

type SegmentBodyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentBodyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentBodyAdvanceRequest) SetAsync(v bool) *SegmentBodyAdvanceRequest {
	s.Async = &v
	return s
}

func (s *SegmentBodyAdvanceRequest) SetReturnForm(v string) *SegmentBodyAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommonImageRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommonImageRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageRequest) SetImageURL(v string) *SegmentCommonImageRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommonImageRequest) SetReturnForm(v string) *SegmentCommonImageRequest {
	s.ReturnForm = &v
	return s
}

type SegmentCommonImageResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SegmentCommonImageResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SegmentCommonImageResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageResponse) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponse) SetRequestId(v string) *SegmentCommonImageResponse {
	s.RequestId = &v
	return s
}

func (s *SegmentCommonImageResponse) SetData(v *SegmentCommonImageResponseData) *SegmentCommonImageResponse {
	s.Data = v
	return s
}

type SegmentCommonImageResponseData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
}

func (s SegmentCommonImageResponseData) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageResponseData) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponseData) SetImageURL(v string) *SegmentCommonImageResponseData {
	s.ImageURL = &v
	return s
}

type SegmentCommonImageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommonImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommonImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentCommonImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentCommonImageAdvanceRequest) SetReturnForm(v string) *SegmentCommonImageAdvanceRequest {
	s.ReturnForm = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imageseg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) SegmentHDSky(request *SegmentHDSkyRequest, runtime *util.RuntimeOptions) (_result *SegmentHDSkyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentHDSkyResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentHDSky"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentHDSkySimply(request *SegmentHDSkyRequest) (_result *SegmentHDSkyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHDSkyResponse{}
	_body, _err := client.SegmentHDSky(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHDSkyAdvance(request *SegmentHDSkyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHDSkyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentHDSkyReq := &SegmentHDSkyRequest{}
	rpcutil.Convert(request, segmentHDSkyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentHDSkyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentHDSkyResp, _err := client.SegmentHDSky(segmentHDSkyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHDSkyResp
	return _result, _err
}

func (client *Client) SegmentHDCommonImage(request *SegmentHDCommonImageRequest, runtime *util.RuntimeOptions) (_result *SegmentHDCommonImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentHDCommonImageResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentHDCommonImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentHDCommonImageSimply(request *SegmentHDCommonImageRequest) (_result *SegmentHDCommonImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHDCommonImageResponse{}
	_body, _err := client.SegmentHDCommonImage(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHDCommonImageAdvance(request *SegmentHDCommonImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHDCommonImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentHDCommonImageReq := &SegmentHDCommonImageRequest{}
	rpcutil.Convert(request, segmentHDCommonImageReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentHDCommonImageReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentHDCommonImageResp, _err := client.SegmentHDCommonImage(segmentHDCommonImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHDCommonImageResp
	return _result, _err
}

func (client *Client) SegmentSkin(request *SegmentSkinRequest, runtime *util.RuntimeOptions) (_result *SegmentSkinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentSkinResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentSkin"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentSkinSimply(request *SegmentSkinRequest) (_result *SegmentSkinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentSkinResponse{}
	_body, _err := client.SegmentSkin(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentSkinAdvance(request *SegmentSkinAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentSkinResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentSkinReq := &SegmentSkinRequest{}
	rpcutil.Convert(request, segmentSkinReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.URLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentSkinReq.URL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentSkinResp, _err := client.SegmentSkin(segmentSkinReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentSkinResp
	return _result, _err
}

func (client *Client) ChangeSky(request *ChangeSkyRequest, runtime *util.RuntimeOptions) (_result *ChangeSkyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ChangeSkyResponse{}
	_body, _err := client.DoRequest(tea.String("ChangeSky"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeSkySimply(request *ChangeSkyRequest) (_result *ChangeSkyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeSkyResponse{}
	_body, _err := client.ChangeSky(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeSkyAdvance(request *ChangeSkyAdvanceRequest, runtime *util.RuntimeOptions) (_result *ChangeSkyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	changeSkyReq := &ChangeSkyRequest{}
	rpcutil.Convert(request, changeSkyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	changeSkyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	changeSkyResp, _err := client.ChangeSky(changeSkyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeSkyResp
	return _result, _err
}

func (client *Client) SegmentLogo(request *SegmentLogoRequest, runtime *util.RuntimeOptions) (_result *SegmentLogoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentLogoResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentLogo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentLogoSimply(request *SegmentLogoRequest) (_result *SegmentLogoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentLogoResponse{}
	_body, _err := client.SegmentLogo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentLogoAdvance(request *SegmentLogoAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentLogoResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentLogoReq := &SegmentLogoRequest{}
	rpcutil.Convert(request, segmentLogoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentLogoReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentLogoResp, _err := client.SegmentLogo(segmentLogoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentLogoResp
	return _result, _err
}

func (client *Client) SegmentScene(request *SegmentSceneRequest, runtime *util.RuntimeOptions) (_result *SegmentSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentSceneResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentScene"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentSceneSimply(request *SegmentSceneRequest) (_result *SegmentSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentSceneResponse{}
	_body, _err := client.SegmentScene(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentSceneAdvance(request *SegmentSceneAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentSceneResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentSceneReq := &SegmentSceneRequest{}
	rpcutil.Convert(request, segmentSceneReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentSceneReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentSceneResp, _err := client.SegmentScene(segmentSceneReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentSceneResp
	return _result, _err
}

func (client *Client) SegmentFood(request *SegmentFoodRequest, runtime *util.RuntimeOptions) (_result *SegmentFoodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentFoodResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentFood"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentFoodSimply(request *SegmentFoodRequest) (_result *SegmentFoodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentFoodResponse{}
	_body, _err := client.SegmentFood(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentFoodAdvance(request *SegmentFoodAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentFoodResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentFoodReq := &SegmentFoodRequest{}
	rpcutil.Convert(request, segmentFoodReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentFoodReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentFoodResp, _err := client.SegmentFood(segmentFoodReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentFoodResp
	return _result, _err
}

func (client *Client) SegmentCloth(request *SegmentClothRequest, runtime *util.RuntimeOptions) (_result *SegmentClothResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentClothResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentCloth"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentClothSimply(request *SegmentClothRequest) (_result *SegmentClothResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentClothResponse{}
	_body, _err := client.SegmentCloth(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentClothAdvance(request *SegmentClothAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentClothResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentClothReq := &SegmentClothRequest{}
	rpcutil.Convert(request, segmentClothReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentClothReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentClothResp, _err := client.SegmentCloth(segmentClothReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentClothResp
	return _result, _err
}

func (client *Client) SegmentAnimal(request *SegmentAnimalRequest, runtime *util.RuntimeOptions) (_result *SegmentAnimalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentAnimalResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentAnimal"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentAnimalSimply(request *SegmentAnimalRequest) (_result *SegmentAnimalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentAnimalResponse{}
	_body, _err := client.SegmentAnimal(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentAnimalAdvance(request *SegmentAnimalAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentAnimalResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentAnimalReq := &SegmentAnimalRequest{}
	rpcutil.Convert(request, segmentAnimalReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentAnimalReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentAnimalResp, _err := client.SegmentAnimal(segmentAnimalReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentAnimalResp
	return _result, _err
}

func (client *Client) SegmentHDBody(request *SegmentHDBodyRequest, runtime *util.RuntimeOptions) (_result *SegmentHDBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentHDBodyResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentHDBody"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentHDBodySimply(request *SegmentHDBodyRequest) (_result *SegmentHDBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHDBodyResponse{}
	_body, _err := client.SegmentHDBody(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHDBodyAdvance(request *SegmentHDBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHDBodyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentHDBodyReq := &SegmentHDBodyRequest{}
	rpcutil.Convert(request, segmentHDBodyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentHDBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentHDBodyResp, _err := client.SegmentHDBody(segmentHDBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHDBodyResp
	return _result, _err
}

func (client *Client) SegmentSky(request *SegmentSkyRequest, runtime *util.RuntimeOptions) (_result *SegmentSkyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentSkyResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentSky"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentSkySimply(request *SegmentSkyRequest) (_result *SegmentSkyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentSkyResponse{}
	_body, _err := client.SegmentSky(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentSkyAdvance(request *SegmentSkyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentSkyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentSkyReq := &SegmentSkyRequest{}
	rpcutil.Convert(request, segmentSkyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentSkyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentSkyResp, _err := client.SegmentSky(segmentSkyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentSkyResp
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.DoRequest(tea.String("GetAsyncJobResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResultSimply(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResult(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentFurniture(request *SegmentFurnitureRequest, runtime *util.RuntimeOptions) (_result *SegmentFurnitureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentFurnitureResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentFurniture"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentFurnitureSimply(request *SegmentFurnitureRequest) (_result *SegmentFurnitureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentFurnitureResponse{}
	_body, _err := client.SegmentFurniture(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentFurnitureAdvance(request *SegmentFurnitureAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentFurnitureResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentFurnitureReq := &SegmentFurnitureRequest{}
	rpcutil.Convert(request, segmentFurnitureReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentFurnitureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentFurnitureResp, _err := client.SegmentFurniture(segmentFurnitureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentFurnitureResp
	return _result, _err
}

func (client *Client) RefineMask(request *RefineMaskRequest, runtime *util.RuntimeOptions) (_result *RefineMaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RefineMaskResponse{}
	_body, _err := client.DoRequest(tea.String("RefineMask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefineMaskSimply(request *RefineMaskRequest) (_result *RefineMaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefineMaskResponse{}
	_body, _err := client.RefineMask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefineMaskAdvance(request *RefineMaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *RefineMaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	refineMaskReq := &RefineMaskRequest{}
	rpcutil.Convert(request, refineMaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	refineMaskReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	refineMaskResp, _err := client.RefineMask(refineMaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = refineMaskResp
	return _result, _err
}

func (client *Client) ParseFace(request *ParseFaceRequest, runtime *util.RuntimeOptions) (_result *ParseFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ParseFaceResponse{}
	_body, _err := client.DoRequest(tea.String("ParseFace"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ParseFaceSimply(request *ParseFaceRequest) (_result *ParseFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ParseFaceResponse{}
	_body, _err := client.ParseFace(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ParseFaceAdvance(request *ParseFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *ParseFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	parseFaceReq := &ParseFaceRequest{}
	rpcutil.Convert(request, parseFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	parseFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	parseFaceResp, _err := client.ParseFace(parseFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = parseFaceResp
	return _result, _err
}

func (client *Client) SegmentVehicle(request *SegmentVehicleRequest, runtime *util.RuntimeOptions) (_result *SegmentVehicleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentVehicleResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentVehicle"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentVehicleSimply(request *SegmentVehicleRequest) (_result *SegmentVehicleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentVehicleResponse{}
	_body, _err := client.SegmentVehicle(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentVehicleAdvance(request *SegmentVehicleAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentVehicleResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentVehicleReq := &SegmentVehicleRequest{}
	rpcutil.Convert(request, segmentVehicleReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentVehicleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentVehicleResp, _err := client.SegmentVehicle(segmentVehicleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentVehicleResp
	return _result, _err
}

func (client *Client) SegmentHair(request *SegmentHairRequest, runtime *util.RuntimeOptions) (_result *SegmentHairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentHairResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentHair"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentHairSimply(request *SegmentHairRequest) (_result *SegmentHairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHairResponse{}
	_body, _err := client.SegmentHair(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHairAdvance(request *SegmentHairAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHairResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentHairReq := &SegmentHairRequest{}
	rpcutil.Convert(request, segmentHairReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentHairReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentHairResp, _err := client.SegmentHair(segmentHairReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHairResp
	return _result, _err
}

func (client *Client) SegmentFace(request *SegmentFaceRequest, runtime *util.RuntimeOptions) (_result *SegmentFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentFaceResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentFace"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentFaceSimply(request *SegmentFaceRequest) (_result *SegmentFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentFaceResponse{}
	_body, _err := client.SegmentFace(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentFaceAdvance(request *SegmentFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentFaceReq := &SegmentFaceRequest{}
	rpcutil.Convert(request, segmentFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentFaceResp, _err := client.SegmentFace(segmentFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentFaceResp
	return _result, _err
}

func (client *Client) SegmentHead(request *SegmentHeadRequest, runtime *util.RuntimeOptions) (_result *SegmentHeadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentHeadResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentHead"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentHeadSimply(request *SegmentHeadRequest) (_result *SegmentHeadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHeadResponse{}
	_body, _err := client.SegmentHead(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHeadAdvance(request *SegmentHeadAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHeadResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentHeadReq := &SegmentHeadRequest{}
	rpcutil.Convert(request, segmentHeadReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentHeadReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentHeadResp, _err := client.SegmentHead(segmentHeadReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHeadResp
	return _result, _err
}

func (client *Client) SegmentCommodity(request *SegmentCommodityRequest, runtime *util.RuntimeOptions) (_result *SegmentCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentCommodityResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentCommodity"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentCommoditySimply(request *SegmentCommodityRequest) (_result *SegmentCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentCommodityResponse{}
	_body, _err := client.SegmentCommodity(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentCommodityAdvance(request *SegmentCommodityAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentCommodityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentCommodityReq := &SegmentCommodityRequest{}
	rpcutil.Convert(request, segmentCommodityReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentCommodityReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentCommodityResp, _err := client.SegmentCommodity(segmentCommodityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentCommodityResp
	return _result, _err
}

func (client *Client) SegmentBody(request *SegmentBodyRequest, runtime *util.RuntimeOptions) (_result *SegmentBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentBodyResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentBody"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentBodySimply(request *SegmentBodyRequest) (_result *SegmentBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentBodyResponse{}
	_body, _err := client.SegmentBody(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentBodyAdvance(request *SegmentBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentBodyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentBodyReq := &SegmentBodyRequest{}
	rpcutil.Convert(request, segmentBodyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentBodyResp, _err := client.SegmentBody(segmentBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentBodyResp
	return _result, _err
}

func (client *Client) SegmentCommonImage(request *SegmentCommonImageRequest, runtime *util.RuntimeOptions) (_result *SegmentCommonImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SegmentCommonImageResponse{}
	_body, _err := client.DoRequest(tea.String("SegmentCommonImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentCommonImageSimply(request *SegmentCommonImageRequest) (_result *SegmentCommonImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentCommonImageResponse{}
	_body, _err := client.SegmentCommonImage(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentCommonImageAdvance(request *SegmentCommonImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentCommonImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageseg"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	segmentCommonImageReq := &SegmentCommonImageRequest{}
	rpcutil.Convert(request, segmentCommonImageReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	segmentCommonImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentCommonImageResp, _err := client.SegmentCommonImage(segmentCommonImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentCommonImageResp
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
